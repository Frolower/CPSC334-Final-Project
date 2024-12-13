package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

// AssignPartToCar inserts a part record for a given chassis_number
func AssignPartToCar(db *sql.DB, chassisNumber string, part *models.Part) error {
	query := `INSERT INTO parts (part_id, quantity, chassis_number) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, part.PartID, part.Quantity, chassisNumber)
	if err != nil {
		log.Printf("Error assigning part: %v", err)
		return err
	}
	return nil
}

// DeletePart deletes a part by part_id
func DeletePart(db *sql.DB, partID string) error {
	query := `DELETE FROM parts WHERE part_id = $1`
	_, err := db.Exec(query, partID)
	return err
}

// GetParts retrieves all parts (no user filtering)
func GetParts(db *sql.DB) ([]models.Part, error) {
	var parts []models.Part
	query := `SELECT part_id, quantity, chassis_number FROM parts`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var part models.Part
		err := rows.Scan(&part.PartID, &part.Quantity, &part.ChassisNumber)
		if err != nil {
			return nil, err
		}
		parts = append(parts, part)
	}
	return parts, nil
}
