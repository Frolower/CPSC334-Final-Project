package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
	"log"
)

func AssignCarToTeam(db *sql.DB, userID, teamID int, car *models.Car) error {
	// Check ownership
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM teams WHERE team_id=$1 AND user_id=$2", teamID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or team not found")
	}

	query := `INSERT INTO cars (chassis_number, make, model, team_id) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(query, car.ChassisNumber, car.Make, car.Model, teamID)
	if err != nil {
		log.Printf("Error inserting car: %v", err)
		return err
	}
	return nil
}

func DeleteCar(db *sql.DB, userID int, chassisNumber string) error {
	// Check ownership
	query := `
		DELETE FROM cars 
		USING teams 
		WHERE cars.team_id = teams.team_id 
		AND teams.user_id = $1 
		AND cars.chassis_number = $2
	`
	_, err := db.Exec(query, userID, chassisNumber)
	return err
}

func GetCarsByUser(db *sql.DB, userID int) ([]models.Car, error) {
	var cars []models.Car
	query := `
		SELECT c.chassis_number, c.make, c.model, c.team_id
		FROM cars c
		JOIN teams t ON c.team_id = t.team_id
		WHERE t.user_id = $1
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ChassisNumber, &car.Make, &car.Model, &car.TeamID)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
