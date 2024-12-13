package models

type Parameter struct {
	Attribute string `json:"attribute"`
	SetupID   int    `json:"setup_id"`
	Value     string `json:"value"`
}
