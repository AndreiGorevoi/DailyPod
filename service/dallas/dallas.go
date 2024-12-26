package dallas

import (
	"DailyPod/config"
	"DailyPod/dto"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const day = time.Hour * 24

type cache struct {
	nextGames struct {
		lastUpdate time.Time
		data       string
	}
}

type Dallas struct {
	config *config.Config
	cache  *cache
}

func NewDallas(cfg *config.Config) *Dallas {
	return &Dallas{
		config: cfg,
		cache:  &cache{},
	}
}

func (dls *Dallas) GetNextGamesStatus() (string, error) {
	if time.Since(dls.cache.nextGames.lastUpdate) < day {
		fmt.Println("Returning from cache")
		return dls.cache.nextGames.data, nil //take data from cache
	}

	games, err := dls.next3Games()
	if err != nil {
		return "", err
	}
	res := formatGamesToString(games)
	dls.cache.nextGames.lastUpdate = time.Now()
	dls.cache.nextGames.data = res
	fmt.Println("Cache is empty")
	return res, nil
}

func (dls *Dallas) next3Games() ([]dto.Game, error) {
	games, err := extractGames()
	filtered := filterGames(games)
	res := make([]dto.Game, 0, 3)
	for i := 0; i < 3 && i < len(filtered); i++ {
		res = append(res, filtered[i])
	}

	return res, err
}

func extractGames() ([]dto.Game, error) {
	f, err := os.Open("service/dallas/mock_data.json")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	var resp dto.GamesResponse

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}

	return resp.Games, nil
}

func filterGames(games []dto.Game) []dto.Game {
	var res []dto.Game

	for _, game := range games {
		if game.Status.Short == 1 {
			res = append(res, game)
		}
	}

	return res
}
