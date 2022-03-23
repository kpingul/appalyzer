package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"net"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"github.com/urfave/cli/v2"
)

var (
	urls []URL
	ports []string 
	ips []IP
)

//Types -- API | CDN
type URL struct {
	Type string 
	URL string
	Path string
}


//Types -- Public | Private
type IP struct {
	Type string 
	IP string 
}

func main() {

	dupURLS := make(map[string]string)

	//init regex 
	regexURL, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	regexIP := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	//Initial CLI App Setup
	app := &cli.App{
		Name:        "Appalyzer",
		Version:     "0.1.0",
		Description: "Static analysis & Threat modeling tool",
		Authors: []*cli.Author{
			{Name: "KP",},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "web", Value: "no", Usage: "Enable web server for GUI", Required: false,},
			&cli.StringFlag{Name: "filepath", Value: "", Usage: "Choose a app project to analyze", Required: false,},
		},
		Action: func(c *cli.Context) error {

			valChecks := false
			webCheck := false

			//input validation checks
		    	if (c.String("web") == "yes" ) {
		    		webCheck = true
		    	} else {

			    	if (c.String("filepath") == "" ) {
			    		valChecks = false
			    	} else {
			    		valChecks = true
			    	}
		    	}

		    	// run if input checks out 
	     		if webCheck {
	     			fmt.Println("RUNNING WEB SERVER")
			    	fileServer := http.FileServer(http.Dir("./frontend")) 
			    	http.Handle("/", fileServer) 
				http.ListenAndServe(":8090", nil)
	     		} else {
			     	if valChecks {

					//walking through project recursively
					err := filepath.Walk(c.String("filepath"),
						func(path string, info os.FileInfo, err error) error {
					    		if err != nil {
					        	return err
					    	}

					    	if !info.IsDir() {

						    	//checking for javascript based files
						    	if filepath.Ext(path) == ".js" || filepath.Ext(path) == ".json" || filepath.Ext(path) == ".ejs" {
						    		file, err := os.Open(path)
								if err != nil {
									log.Fatal(err)
								}
								defer file.Close()

								scanner := bufio.NewScanner(file)
								buf := make([]byte, 0, 64*1024)
								scanner.Buffer(buf, 1024*1024)
								
								for scanner.Scan() {

									//test case #1 HTTP/HTTPS
									if regexURL.MatchString(scanner.Text()) {
										/*
										CDN Detection
										-link contains href which maps to URL 
											-if line contains link and href, this will have a URL attached to it 
										-script contains src which maps to URL
											-if line contains script and src, this will have a URL attached to it 
										
										if scanner.Text() contains link + href= 
											CSS file
										if scanner.Text() contains script + src=
											JS file
										*/
										if strings.Contains(scanner.Text(), "link") && strings.Contains(scanner.Text(), "href=") {
											if _, ok := dupURLS[regexURL.FindString(scanner.Text())]; !ok {
												
												dupURLS[regexURL.FindString(scanner.Text())] = regexURL.FindString(scanner.Text()) 
											
												urls = append(urls, URL{
													Type: "CDN",
													URL: regexURL.FindString(scanner.Text()),
													Path: path,
												})

											}
										}
										if strings.Contains(scanner.Text(), "script") && strings.Contains(scanner.Text(), "src=") {
											if _, ok := dupURLS[regexURL.FindString(scanner.Text())]; !ok {
												
												dupURLS[regexURL.FindString(scanner.Text())] = regexURL.FindString(scanner.Text()) 
											
												urls = append(urls, URL{
													Type: "CDN",
													URL: regexURL.FindString(scanner.Text()),
													Path: path,
												})

											}
										}

									}

									//test case #2 IP Addresses
									if regexIP.MatchString(scanner.Text()) {
										ipCheck := checkIPAddress(regexIP.FindString(scanner.Text()))

										/*
										IP detection filtering
											-need to find a way to filter non IP addresses

										*/


										if ipCheck {
											ip := regexIP.FindString(scanner.Text())
											ips - append(ips, IP{
												Type: privateIPCheck(ip),
												IP: ip,
											})
										}
									}
								}

								if err := scanner.Err(); err != nil {
									log.Fatal(err)
								}
							}


					    	}


					    	return nil
					})

					for i := 0; i < len(urls); i++ {
						fmt.Println(urls[i])
					}
					if err != nil {
					    log.Println(err)
					}


			     	} else {
			     		fmt.Println("stop program..")
			     		return nil
			     	}
	     		}



		     	return nil
	    	},
	}



	//Run CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/* Utility */

// Check if a ip is private.
func privateIPCheck(ip string) string {
    	ipAddress := net.ParseIP(ip)

    	if ip == "127.0.0.1" {
    		return "Private"
    	} else if ipAddress.IsPrivate() {
    		return "Private"
    	} else {
    		return "Public"
    	}
}

// check ip address validity 
func checkIPAddress(ip string) bool {
    	if net.ParseIP(ip) == nil {
    		return false
    	} else {
    		return true
    	}
}