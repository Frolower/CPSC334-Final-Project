package models

type Stage struct {
	StageID        int    `json:"stage_id"`
	StageNumber    int    `json:"stage_number"`
	ChampionshipID int    `json:"championship_id"`
	Track          string `json:"track"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
}
