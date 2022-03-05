package main

import (
	"log"
	"os"
	"fmt"
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
			err := filepath.Walk(`.`,
				func(path string, info os.FileInfo, err error) error {
			    		if err != nil {
			        	return err
			    	}
			    	fmt.Println(path, info.Size())
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