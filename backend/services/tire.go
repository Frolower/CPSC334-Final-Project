package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

// AssignTireToCar inserts a tire for a given chassis_number
func AssignTireToCar(db *sql.DB, chassisNumber string, tire *models.Tire) error {
	query := `INSERT INTO tires (tire_id, tread_remaining, compound, chassis_number) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, tire.TireID, tire.TreadRemaining, tire.Compound, chassisNumber)
	if err != nil {
		log.Printf("Error assigning tire: %v", err)
		return err
	}
	return nil
}

// DeleteTire deletes a tire by tire_id
func DeleteTire(db *sql.DB, tireID string) error {
	query := `DELETE FROM tires WHERE tire_id = $1`
	_, err := db.Exec(query, tireID)
	return err
}

// GetTiresByChassisNumber retrieves all tires for a given chassis_number
func GetTiresByChassisNumber(db *sql.DB, chassisNumber string) ([]models.Tire, error) {
	var tires []models.Tire
	query := `SELECT tire_id, tread_remaining, compound, chassis_number FROM tires WHERE chassis_number = $1`
	rows, err := db.Query(query, chassisNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Tire
		if err := rows.Scan(&t.TireID, &t.TreadRemaining, &t.Compound, &t.ChassisNumber); err != nil {
			return nil, err
		}
		tires = append(tires, t)
	}
	return tires, nil
}

// UpdateTire updates a tires details by tire_id
func UpdateTire(db *sql.DB, tireID string, tire *models.Tire) error {
	query := `UPDATE tires SET tread_remaining = $1, compound = $2, chassis_number = $3 WHERE tire_id = $4`
	_, err := db.Exec(query, tire.TreadRemaining, tire.Compound, tire.ChassisNumber, tireID)
	return err
}
