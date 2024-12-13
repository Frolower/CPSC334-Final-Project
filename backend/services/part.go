package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

// AssignPartToCar inserts a part record for a given chassis_number
func AssignPartToCar(db *sql.DB, chassisNumber string, part *models.Part) error {
	query := `INSERT INTO parts (part_id, part_name, quantity, chassis_number) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, part.PartID, part.PartName, part.Quantity, chassisNumber)
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

// GetPartsByChassisNumber retrieves all parts associated with the given chassis_number
func GetPartsByChassisNumber(db *sql.DB, chassisNumber string) ([]models.Part, error) {
	var parts []models.Part
	query := `SELECT part_id, part_name, quantity, chassis_number FROM parts WHERE chassis_number = $1`
	rows, err := db.Query(query, chassisNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var part models.Part
		if err := rows.Scan(&part.PartID, &part.PartName, &part.Quantity, &part.ChassisNumber); err != nil {
			return nil, err
		}
		parts = append(parts, part)
	}
	return parts, nil
}

// UpdatePart updates the part details by part_id
func UpdatePart(db *sql.DB, partID string, part *models.Part) error {
	query := `UPDATE parts SET part_name=$1, quantity=$2, chassis_number=$3 WHERE part_id=$4`
	_, err := db.Exec(query, part.PartName, part.Quantity, part.ChassisNumber, partID)
	if err != nil {
		log.Printf("Error updating part %s: %v", partID, err)
		return err
	}
	return nil
}
