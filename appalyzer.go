package main

import (
	"log"
	"os"
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

		     	return nil
	    	},
	}



	//Run CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}