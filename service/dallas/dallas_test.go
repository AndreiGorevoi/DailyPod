package dallas

import (
	"DailyPod/dto"
	"testing"
)

func TestFilterGames(t *testing.T) {
	games := []dto.Game{
		{
			Status: dto.Status{
				Short: 1,
			},
		},
		{
			Status: dto.Status{
				Short: 2,
			},
		},
		{
			Status: dto.Status{
				Short: 1,
			},
		},
	}

	res := filterGames(games)

	for _, v := range res {
		if v.Status.Short != 1 {
			t.Error("Games with short 1 should left after filtering")
		}
	}

	if len(res) != 2 {
		t.Errorf("Expected 2 games were left, was: %d", len(res))
	}

}
