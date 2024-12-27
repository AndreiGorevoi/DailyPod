package dallas

import (
	"DailyPod/dto"
	"testing"
	"time"
)

func TestFormatGamesToString(t *testing.T) {
	startTime, err := time.Parse(time.RFC3339, "2024-12-25T12:00:05Z")
	if err != nil {
		t.Fatalf("Failed to parse time: %v", err)
	}
	data := map[string]struct {
		input    []dto.Game
		expected string
	}{
		`one games`: {
			input: []dto.Game{
				{
					Date: dto.Date{
						Start: startTime,
					},
					Teams: dto.Teams{
						Visitors: dto.Team{
							Nickname: "Dallas",
						},
						Home: dto.Team{
							Nickname: "LosAngeles",
						},
					},
				},
			},
			expected: "1: LosAngeles - Dallas. 🗓️ Wed,Dec 25 - 1:00PM(CET)\n",
		},
		`two games`: {
			input: []dto.Game{
				{
					Date: dto.Date{
						Start: startTime.Add(time.Hour * 24),
					},
					Teams: dto.Teams{
						Visitors: dto.Team{
							Nickname: "Portland",
						},
						Home: dto.Team{
							Nickname: "GSW",
						},
					},
				},
				{
					Date: dto.Date{
						Start: startTime,
					},
					Teams: dto.Teams{
						Visitors: dto.Team{
							Nickname: "Dallas",
						},
						Home: dto.Team{
							Nickname: "LosAngeles",
						},
					},
				},
			},
			expected: "1: GSW - Portland. 🗓️ Thu,Dec 26 - 1:00PM(CET)\n2: LosAngeles - Dallas. 🗓️ Wed,Dec 25 - 1:00PM(CET)\n",
		},
		`empty input`: {
			input:    nil,
			expected: "",
		},
	}

	for name, tt := range data {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := formatGamesToString(tt.input)
			if tt.expected != got {
				t.Errorf("got: %s. expected: %s", got, tt.expected)
			}
		})
	}
}
