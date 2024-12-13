package models

type Part struct {
	PartID        string `json:"part_id"`
	Quantity      int    `json:"quantity"`
	ChassisNumber string `json:"chassis_number"`
}
