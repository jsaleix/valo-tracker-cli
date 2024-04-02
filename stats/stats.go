package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"valo_tracker/config"

	gcl "github.com/MaphicalYng/golang-cmd-loading"
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

func fetchHistoryOfPlayer() (data []PLAYER_STATS, err error) {
	//Fetching data
	fullUrl := fmt.Sprintf("%s/v3/matches/eu/%s/%s", config.API_ENDPOINT, config.USERNAME, config.TAG)
	var res *http.Response

	gcl.WithLoadingMessage(func(cancelLoading func()) {
		res, err = http.Get(fullUrl)
		cancelLoading()
	}, "Fetching...")

	if err != nil || !strings.Contains(res.Status, "200") {
		// log.Fatalln(err)
		return data, errors.New("error while fetching data")
	}
	defer res.Body.Close()

	// decoding result
	var cResp API_RESULT

	if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
		log.Println(err)
		return data, err
		// log.Fatal("An error occurred, please try again mate")
	}

	for _, v := range cResp.Data {
		for _, p := range v.Players.AllPlayers {
			if p.Name == config.USERNAME && p.Tag == config.TAG {
				data = append(data, p.Stats)
			}
		}
	}

	return data, nil

}

func displayHistory(p []PLAYER_STATS) {
	if len(p) == 0 {
		fmt.Printf("There is no game to display")
		return
	}

	for _, v := range p {
		fmt.Printf("KDA: %d/%d/%d\n", v.Kills, v.Deaths, v.Assists)
	}
}

func Run() {
	res, err := fetchHistoryOfPlayer()
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("#### VALORANT GAME HISTORY ####\n")
		displayHistory(res)
		fmt.Printf("#### END OF HISTORY ####\n")
	}
}
