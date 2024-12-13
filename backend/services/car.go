package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

// AssignCarToTeam inserts a car record for the given team without ownership checks
func AssignCarToTeam(db *sql.DB, teamID int, car *models.Car) error {
	query := `INSERT INTO cars (chassis_number, make, model, team_id) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, car.ChassisNumber, car.Make, car.Model, teamID)
	if err != nil {
		log.Printf("Error inserting car: %v", err)
		return err
	}
	return nil
}

// DeleteCar deletes a car by chassis_number without ownership checks
func DeleteCar(db *sql.DB, chassisNumber string) error {
	query := `DELETE FROM cars WHERE chassis_number = $1`
	_, err := db.Exec(query, chassisNumber)
	return err
}

// GetCars retrieves all cars (no user filtering)
func GetCars(db *sql.DB) ([]models.Car, error) {
	var cars []models.Car
	query := `SELECT chassis_number, make, model, team_id FROM cars`
	rows, err := db.Query(query)
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
