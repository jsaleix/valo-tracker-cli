package main

import (
	"flag"
	"fmt"
	"valo_tracker/config"
	"valo_tracker/stats"
)

func main() {
	resetCreds := flag.Bool("reset", false, "Redefine your tag and username")
	helpCmd := flag.Bool("help", false, "Show the available parameters")
	currentCmd := flag.Bool("current", false, "Display your current tag and username")

	flag.Parse()
	config.Init()

	if *helpCmd {
		fmt.Printf("--current: Display your current tag and username\n")
		fmt.Printf("--reset: Redefine your tag and username\n")
		fmt.Printf("--help: Show the available parameters\n")
		return
	}

	if *currentCmd {
		fmt.Printf("Current: %s #%s\n", config.USERNAME, config.TAG)
		return
	}

	if *resetCreds {
		config.PromptUser()
	}

	stats.Run()
}
