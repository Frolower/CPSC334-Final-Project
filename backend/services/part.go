// services/part.go
package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
	"log"
)

func AssignPartToCar(db *sql.DB, userID int, chassisNumber string, part *models.Part) error {
	queryOwnership := `
        SELECT COUNT(*) 
        FROM cars c 
        JOIN teams t ON c.team_id = t.team_id 
        WHERE c.chassis_number = $1 AND t.user_id = $2
    `
	var count int
	if err := db.QueryRow(queryOwnership, chassisNumber, userID).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or car not found")
	}

	query := `INSERT INTO parts (part_id, quantity, chassis_number) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, part.PartID, part.Quantity, chassisNumber)
	if err != nil {
		log.Printf("Error assigning part: %v", err)
		return err
	}
	return nil
}

func DeletePart(db *sql.DB, userID int, partID string) error {
	query := `
		DELETE FROM parts 
		USING cars, teams 
		WHERE parts.chassis_number = cars.chassis_number 
		AND cars.team_id = teams.team_id
		AND teams.user_id = $1
		AND parts.part_id = $2
	`
	_, err := db.Exec(query, userID, partID)
	return err
}

func GetPartsByUser(db *sql.DB, userID int) ([]models.Part, error) {
	var parts []models.Part
	query := `
		SELECT p.part_id, p.quantity, p.chassis_number
		FROM parts p
		JOIN cars c ON p.chassis_number = c.chassis_number
		JOIN teams t ON c.team_id = t.team_id
		WHERE t.user_id = $1
	`
	rows, err := db.Query(query, userID)
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
