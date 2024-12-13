package models

type Session struct {
	SessionID   int    `json:"session_id"`
	StageID     int    `json:"stage_id"`
	Type        string `json:"type"`
	SessionDate string `json:"session_date"`
	StartTime   string `json:"start_time"`
	Weather     string `json:"weather"`
	Temperature int    `json:"temperature"`
	Humidity    int    `json:"humidity"`
}
