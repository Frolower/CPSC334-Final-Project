package models

type Car struct {
	ChassisNumber string `json:"chassis_number"`
	Make          string `json:"make"`
	Model         string `json:"model"`
	TeamID        int    `json:"team_id"`
}
