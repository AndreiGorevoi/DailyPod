package dto

import "time"

type GamesResponse struct {
	Response []Response `json:"response"`
}

type Response struct {
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
