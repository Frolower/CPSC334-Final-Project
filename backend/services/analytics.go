package services

import (
	"database/sql"
)

// GetFastestLap returns the lap_number and lap_time of the fastest lap for a given session_id
func GetFastestLap(db *sql.DB, sessionID int) (int, string, error) {
	query := `
		SELECT lap_number, lap_time
		FROM laps
		WHERE session_id = $1
		ORDER BY lap_time ASC
		LIMIT 1
	`
	var lapNumber int
	var lapTime string
	err := db.QueryRow(query, sessionID).Scan(&lapNumber, &lapTime)
	if err != nil {
		return 0, "", err
	}
	return lapNumber, lapTime, nil
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
