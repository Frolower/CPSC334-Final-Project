package models

type Lap struct {
	LapNumber int    `json:"lap_number"`
	SessionID int    `json:"session_id"`
	LapTime   string `json:"lap_time"`
}
