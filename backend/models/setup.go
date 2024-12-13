package models

type Setup struct {
	SetupID        int    `json:"setup_id"`
	SetupName      string `json:"setup_name"`
	TrackName      string `json:"track_name"`
	ChassisNumber  string `json:"chassis_number"`
	ChampionshipID int    `json:"championship_id"`
}
