package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"regexp"
	"github.com/urfave/cli/v2"
)

var (
	urls []string
	ports []string 
)

func main() {

	//init regex 
	regexURL, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)


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


			//walking through project recursively
			err := filepath.Walk(`.`,
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
					}

					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}
				}
			    	if filepath.Ext(path) == ".json" {
			    	
				
				}
			    	if filepath.Ext(path) == ".ejs" {
				
				}
			    	

			    	return nil
			})
			
			if err != nil {
			    log.Println(err)
			}

			fmt.Println(len(urls))

		     	return nil
	    	},
	}



	//Run CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}