package stats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"valo_tracker/config"
)

type PLAYER_STATS struct {
	Score   int `json:"score"`
	Kills   int `json:"kills"`
	Deaths  int `json:"deaths"`
	Assists int `json:"assists"`
}

type PLAYER_DATA struct {
	Name  string       `json:"name"`
	Tag   string       `json:"tag"`
	Stats PLAYER_STATS `json:"stats"`
}

type GAME_PLAYERS struct {
	AllPlayers []PLAYER_DATA `json:"all_players"`
}

type GAME struct {
	Players GAME_PLAYERS `json:"players"`
}

type API_RESULT struct {
	Status int    `json:"status"`
	Data   []GAME `json:"data"`
}

func FetchHistoryOfPlayer() (data []PLAYER_STATS, ok bool) {
	fullUrl := fmt.Sprintf("%s/v3/matches/eu/%s/%s", config.API_ENDPOINT, config.USERNAME, config.TAG)
	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln(err)
		return data, false
	}
	defer res.Body.Close()

	var cResp API_RESULT

	if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
		log.Println(err)
		log.Fatal("An error occurred, please try again mate")
	}

	for _, v := range cResp.Data {
		for _, p := range v.Players.AllPlayers {
			if p.Name == config.USERNAME && p.Tag == config.TAG {
				data = append(data, p.Stats)
			}
		}
	}

	return data, true
}
