package services

import (
	"Ariadne_Management/models"
	"database/sql"
)

// CreateLap inserts a new lap record
func CreateLap(db *sql.DB, lap *models.Lap) error {
	query := `INSERT INTO laps (lap_number, session_id, lap_time) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, lap.LapNumber, lap.SessionID, lap.LapTime)
	return err
}

// GetLapsBySessionID retrieves all laps for a given session_id
func GetLapsBySessionID(db *sql.DB, sessionID int) ([]models.Lap, error) {
	var laps []models.Lap
	query := `SELECT lap_number, lap_time FROM laps WHERE session_id = $1`
	rows, err := db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lap models.Lap
		if err := rows.Scan(&lap.LapNumber, &lap.LapTime); err != nil {
			return nil, err
		}
		lap.SessionID = sessionID
		laps = append(laps, lap)
	}
	return laps, nil
}

// UpdateLap updates a laps time
func UpdateLap(db *sql.DB, sessionID, lapNumber int, newLapTime string) error {
	query := `UPDATE laps SET lap_time=$1 WHERE session_id=$2 AND lap_number=$3`
	_, err := db.Exec(query, newLapTime, sessionID, lapNumber)
	return err
}

// DeleteLap deletes a lap
func DeleteLap(db *sql.DB, sessionID, lapNumber int) error {
	query := `DELETE FROM laps WHERE session_id=$1 AND lap_number=$2`
	_, err := db.Exec(query, sessionID, lapNumber)
	return err
}
