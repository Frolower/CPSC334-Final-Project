package models

type Pilot struct {
	DocumentNumber int    `json:"document_number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Position       string `json:"position"`
	TeamID         int    `json:"team_id"`
}
