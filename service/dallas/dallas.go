package dallas

import (
	"DailyPod/config"
	"DailyPod/dto"
	"encoding/json"
	"io"
	"os"
)

type Dallas struct {
	config *config.Config
}

func NewDallas(cfg *config.Config) *Dallas {
	return &Dallas{
		config: cfg,
	}
}

func (dls *Dallas) Next3Games() (res [3]dto.Game, err error) {
	games, err := extractGames()
	filtered := filterGames(games)
	for i := 0; i < 3 && i < len(filtered); i++ {
		res[i] = filtered[i]
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
