package config

import "fmt"

const (
	RESULT_LIMIT = 0
	API_ENDPOINT = "https://api.henrikdev.xyz/valorant"
)

// var USERNAME = "impirashield"
// var TAG = "6274"

var USERNAME string
var TAG string

func Init() {
	if TAG == "" || USERNAME == "" {
		promptUser()
	}
}

func promptUser() {
	fmt.Printf("You have not set your Valorant Tag, please specify it \n")

	var tag string
	var username string

	fmt.Print("\tTag(e.g. 6274): ")
	for tag == "" || len(tag) > 4 {
		fmt.Scanf("%s", &tag)
	}

	fmt.Print("\tUsername: ")
	for username == "" {
		fmt.Scanf("%s", &username)
	}

	USERNAME = username
	TAG = tag
}
