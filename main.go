package main

import (
	"valo_tracker/config"
	"valo_tracker/stats"
)

func main() {
	config.Init()
	stats.Run()
}
