package models

type Tire struct {
	TireID         string  `json:"tire_id"`
	TreadRemaining float64 `json:"tread_remaining"`
	Compound       string  `json:"compound"`
	ChassisNumber  string  `json:"chassis_number"`
}
