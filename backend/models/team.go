// /models/user.go

package models

// User model represents a user in the system
type Team struct {
	Team_ID   int    `json:"team_id"`
	User_ID   string `json:"user_id"`
	Team_name string `json:"team_name"`
}
