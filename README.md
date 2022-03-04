# appalyzer

Static code analysis tool + threat modeling tool

## Purpose

The goal of this project is to automate basic threat modeling for codebases. The project will label the potential attack surfaces that I call indicators. They can be network based, credentials, or keys. Once I gather the data, these indicators will be mapped out graphically through an interface that you can interact with. This is just a POC that I wanted to create for fun to learn more about threat modeling and static code analysis. 

![Untitled](https://user-images.githubusercontent.com/11414669/156837444-4a2fc0c7-5dcb-4d47-80d4-e9c9586e14f2.png)

## Features

* Static code analysis 

## Language and File Support

Javacript (js, json for now..)

## Indicators

* Network (case insenstive )
	- HTTP/HTTPS
	- port
* Credentials (case insenstive )
	- username
	- password
	- name
	- lastname
* Keys
	- key

## Data Model

* Type - indicator type
* Finding - indicator
* Location - absolute path + file name

## Sources

https://cheatsheetseries.owasp.org/cheatsheets/Threat_Modeling_Cheat_Sheet.html
https://www.npmjs.com/package/common-js-file-extensions




