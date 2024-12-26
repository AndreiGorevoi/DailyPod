package dallas

import (
	"DailyPod/config"
	"DailyPod/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	day                   = time.Hour * 24
	dallasMavericksTeamId = 8
)

type cache struct {
	nextGames struct {
		lastUpdate time.Time
		data       string
	}
}

type Dallas struct {
	config *config.Config
	cache  *cache
	client *http.Client
}

func NewDallas(cfg *config.Config, client *http.Client) *Dallas {
	return &Dallas{
		config: cfg,
		cache:  &cache{},
		client: client,
	}
}

func (dls *Dallas) GetNextGamesStatus() (string, error) {
	if time.Since(dls.cache.nextGames.lastUpdate) < day {
		return dls.cache.nextGames.data, nil //take data from cache
	}

	games, err := dls.next3Games()
	if err != nil {
		return "", err
	}
	res := formatGamesToString(games)
	dls.cache.nextGames.lastUpdate = time.Now()
	dls.cache.nextGames.data = res
	return res, nil
}

func (dls *Dallas) next3Games() (games []dto.Game, err error) {
	if dls.config.MockDbData {
		games, err = extractGamesFromMock()
	} else {
		games, err = dls.callApiNba()
	}

	if err != nil {
		return nil, err
	}

	filtered := filterGames(games)
	res := make([]dto.Game, 0, 3)
	for i := 0; i < 3 && i < len(filtered); i++ {
		res = append(res, filtered[i])
	}

	return res, err
}

func extractGamesFromMock() ([]dto.Game, error) {
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

func (dls *Dallas) callApiNba() ([]dto.Game, error) {
	url := dls.buildUrl()

	resp, err := dls.makeRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return dls.parseResponse(resp)
}

func (dls *Dallas) buildUrl() string {
	//be carefull about season, not sure if it will work with currect year
	season := time.Now().Year()
	return fmt.Sprintf("%s/games?season=%d&team=%d", dls.config.ApiNbaUrl, season, dallasMavericksTeamId)
}

func (dls *Dallas) makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
	req.Header.Set("x-rapidapi-key", dls.config.NbaApiKey)

	return dls.client.Do(req)
}

func (dls *Dallas) parseResponse(resp *http.Response) ([]dto.Game, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gameResponse dto.GamesResponse
	err = json.Unmarshal(data, &gameResponse)
	if err != nil {
		return nil, err
	}

	return gameResponse.Games, nil
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
