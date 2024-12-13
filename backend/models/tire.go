package models

type Tire struct {
	TireID         string `json:"tire_id"`
	TreadRemaining int    `json:"tread_remaining"`
	Compound       string `json:"compound"`
	ChassisNumber  string `json:"chassis_number"`
}
