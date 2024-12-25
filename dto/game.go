package dto

import "time"

type GamesResponse struct {
	Games []Game `json:"response"`
}

type Game struct {
	Date   Date   `json:"date"`
	Status Status `json:"status"`
	Teams  Teams  `json:"teams"`
}

type Date struct {
	Start time.Time `json:"start"`
}

type Status struct {
	Short byte   `json:"short"`
	Long  string `json:"long"`
}

type Teams struct {
	Visitors Team `json:"visitors"`
	Home     Team `json:"home"`
}

type Team struct {
	Name string `json:"name"`
}
