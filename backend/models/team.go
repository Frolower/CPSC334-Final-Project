package models

// Team model represents a team in the system
type Team struct {
	Team_ID   int    `json:"team_id"`
	User_ID   int    `json:"user_id"`
	Team_name string `json:"team_name"`
}
