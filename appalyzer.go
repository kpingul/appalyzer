package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"github.com/urfave/cli/v2"
)

func main() {
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
			err := filepath.Walk(`C:\Users\14152\Downloads\Eat36Five-master\Eat36Five-master`,
				func(path string, info os.FileInfo, err error) error {
			    		if err != nil {
			        	return err
			    	}

			    	//checking for javascript based files
			    	if filepath.Ext(path) == ".js" {
			    		fmt.Println(path)
			    		file, err := os.Open(path)
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					scanner := bufio.NewScanner(file)
					
					for scanner.Scan() {
						// fmt.Println(scanner.Text())
					}

					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}
				}
			    	if filepath.Ext(path) == ".json" {
			    		fmt.Println(path)
				
				}
			    	if filepath.Ext(path) == ".ejs" {
			    		fmt.Println(path)
				
				}
			    	

			    	return nil
			})
			
			if err != nil {
			    log.Println(err)
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