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

func (dls *Dallas) GetNextGamesStatus() (string, error) {
	games, err := dls.next3Games()
	if err != nil {
		return "", err
	}
	res := formatGamesToString(games)
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
