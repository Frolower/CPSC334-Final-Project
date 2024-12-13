package services

import (
	"database/sql"
	"log"
)

// GetFastestLap returns the lap_number and lap_time of the fastest lap for a given session_id
func GetFastestLap(db *sql.DB, sessionID int) (int, string, float64, error) {
	query := `
		SELECT lap_number, lap_time, EXTRACT(EPOCH FROM lap_time) AS lap_seconds
		FROM laps
		WHERE session_id = $1
		ORDER BY lap_time ASC
		LIMIT 1
	`
	var lapNumber int
	var lapTime string
	var lapSeconds float64
	err := db.QueryRow(query, sessionID).Scan(&lapNumber, &lapTime, &lapSeconds)
	if err != nil {
		return 0, "", 0, err
	}
	return lapNumber, lapTime, lapSeconds, nil
}

// GetAverageLapTime returns the average lap time (in seconds) for a given session_id
func GetAverageLapTime(db *sql.DB, sessionID int) (float64, error) {
	query := `
		SELECT AVG(EXTRACT(EPOCH FROM lap_time))
		FROM laps
		WHERE session_id = $1
	`
	var avgSeconds float64
	err := db.QueryRow(query, sessionID).Scan(&avgSeconds)
	if err != nil {
		return 0, err
	}
	return avgSeconds, nil
}

// GetPartsCountForCar gets the number of parts assigned ofr a car
func GetPartsCountForCar(db *sql.DB, chassisNumber string) (int, error) {
	query := `SELECT COUNT(*) FROM parts WHERE chassis_number = $1`
	var count int
	err := db.QueryRow(query, chassisNumber).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetTireCountForCar returns the number of tires for a given chassis_number
func GetTireCountForCar(db *sql.DB, chassisNumber string) (int, error) {
	query := `SELECT COUNT(*) FROM tires WHERE chassis_number = $1`
	var count int
	err := db.QueryRow(query, chassisNumber).Scan(&count)
	if err != nil {
		log.Printf("Error fetching tire count for car %s: %v", chassisNumber, err)
		return 0, err
	}
	return count, nil
}

// GetAverageTreadForCar returns the average tread remaining for a given chassis_number
func GetAverageTreadForCar(db *sql.DB, chassisNumber string) (float64, error) {
	query := `SELECT AVG(tread_remaining) FROM tires WHERE chassis_number = $1`
	var avg float64
	err := db.QueryRow(query, chassisNumber).Scan(&avg)
	if err != nil {
		log.Printf("Error fetching average tread for car %s: %v", chassisNumber, err)
		return 0, err
	}
	return avg, nil
}

// GetTiresCountByTreadAndCompoundAndCar returns the count of tires with tread >= x, compound = y for a specific car
func GetTiresCountByTreadAndCompoundAndCar(db *sql.DB, chassisNumber string, x float64, compound string) (int, error) {
	query := `SELECT COUNT(*) FROM tires WHERE chassis_number = $1 AND tread_remaining >= $2 AND compound = $3`
	var count int
	err := db.QueryRow(query, chassisNumber, x, compound).Scan(&count)
	if err != nil {
		log.Printf("Error fetching tires count for car %s with tread >= %.1f and compound %s: %v", chassisNumber, x, compound, err)
		return 0, err
	}
	return count, nil
}
