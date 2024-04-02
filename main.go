package main

import (
	"fmt"
	"valo_tracker/stats"
)

func displayHistory(p []stats.PLAYER_STATS) {
	for _, v := range p {
		fmt.Printf("KDA: %d/%d/%d\n", v.Kills, v.Deaths, v.Assists)
	}
}
func main() {
	res, ok := stats.FetchHistoryOfPlayer()
	if ok {
		fmt.Printf("#### VALORANT GAME HISTORY ####\n")
		displayHistory(res)
		fmt.Printf("#### END OF HISTORY ####\n")
	}
	// name := flag.String("name", "world", "The name to greet.")
	// flag.Parse()
	// fmt.Printf("Hello, %s!\n", *name)
}
