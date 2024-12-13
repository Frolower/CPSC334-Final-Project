package models

type Championship struct {
	ChampionshipID   int    `json:"championship_id"`
	TeamID           int    `json:"team_id"`
	ChampionshipName string `json:"championship_name"`
	TeamStandings    int    `json:"team_standings"`
}
