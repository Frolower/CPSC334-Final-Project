package models

type Part struct {
	PartID        string `json:"part_id"`
	PartName      string `json:"part_name"`
	Quantity      int    `json:"quantity"`
	ChassisNumber string `json:"chassis_number"`
}
