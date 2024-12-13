package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

// AssignCarToTeam inserts a car record for the given team
func AssignCarToTeam(db *sql.DB, teamID int, car *models.Car) error {
	query := `INSERT INTO cars (chassis_number, make, model, team_id) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, car.ChassisNumber, car.Make, car.Model, teamID)
	if err != nil {
		log.Printf("Error inserting car: %v", err)
		return err
	}
	return nil
}

// DeleteCar deletes a car by chassis_number
func DeleteCar(db *sql.DB, chassisNumber string) error {
	query := `DELETE FROM cars WHERE chassis_number = $1`
	_, err := db.Exec(query, chassisNumber)
	return err
}

// GetCarsByTeamID retrieves all cars belonging to a specific team
func GetCarsByTeamID(db *sql.DB, teamID int) ([]models.Car, error) {
	var cars []models.Car
	query := `SELECT chassis_number, make, model, team_id FROM cars WHERE team_id = $1`
	rows, err := db.Query(query, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ChassisNumber, &car.Make, &car.Model, &car.TeamID); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

// UpdateCar updates the car details by chassis_number
func UpdateCar(db *sql.DB, chassisNumber string, car *models.Car) error {
	query := `UPDATE cars SET make=$1, model=$2, team_id=$3 WHERE chassis_number=$4`
	_, err := db.Exec(query, car.Make, car.Model, car.TeamID, chassisNumber)
	return err
}
