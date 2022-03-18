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
	"github.com/urfave/cli/v2"
)

var (
	urls []string
	ports []string 
)

//Types -- API | CDN
type URL struct {
	Type string 
	URL string
}


//Types -- Public | Private
type IP struct {
	Type string 
	IP string 
}

func main() {

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

					    	//checking for javascript based files
					    	if filepath.Ext(path) == ".js" {
					    		file, err := os.Open(path)
							if err != nil {
								log.Fatal(err)
							}
							defer file.Close()

							scanner := bufio.NewScanner(file)
							
							for scanner.Scan() {

								//test case #1 HTTP/HTTPS
								if regexURL.MatchString(scanner.Text()) {
									url := regexURL.FindString(scanner.Text())
									urls = append(urls, url)
								}

								//test case #2 IP Addresses
								if regexIP.MatchString(scanner.Text()) {
									ipCheck := checkIPAddress(regexIP.FindString(scanner.Text()))

									if ipCheck {
										ip := regexIP.FindString(scanner.Text())
										fmt.Println(ip)
									}
								}
							}

							if err := scanner.Err(); err != nil {
								log.Fatal(err)
							}
						}
					    	if filepath.Ext(path) == ".json" {
					    	
						
						}
					    	if filepath.Ext(path) == ".ejs" {
						
						}

			     			if webCheck {
				     			//setup http web server and API's
						    	fileServer := http.FileServer(http.Dir("./frontend")) 
						    	http.Handle("/", fileServer) 
							http.ListenAndServe(":8090", nil)
						}
					    	

					    	return nil
					})
					
					if err != nil {
					    log.Println(err)
					}

					fmt.Println(len(urls))

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

    	if ipAddress.IsPrivate() {
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