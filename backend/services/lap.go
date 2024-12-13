package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
)

// CreateLap adds a lap to a session
func CreateLap(db *sql.DB, userID int, lap *models.Lap) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM sessions se
		JOIN stages s ON se.stage_id=s.stage_id
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id=t.team_id
		WHERE se.session_id=$1 AND t.user_id=$2
	`, lap.SessionID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or session not found")
	}

	query := `INSERT INTO laps (lap_number, session_id, lap_time) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, lap.LapNumber, lap.SessionID, lap.LapTime)
	return err
}

// GetLapsByUser retrieves all laps for sessions owned by the user
func GetLapsByUser(db *sql.DB, userID int) ([]models.Lap, error) {
	var laps []models.Lap
	query := `
		SELECT l.lap_number, l.session_id, l.lap_time
		FROM laps l
		JOIN sessions se ON l.session_id=se.session_id
		JOIN stages s ON se.stage_id=s.stage_id
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id=t.team_id
		WHERE t.user_id=$1
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lap models.Lap
		if err := rows.Scan(&lap.LapNumber, &lap.SessionID, &lap.LapTime); err != nil {
			return nil, err
		}
		laps = append(laps, lap)
	}
	return laps, nil
}

// GetLapByKey retrieves a single lap by session_id and lap_number
func GetLapByKey(db *sql.DB, userID, sessionID, lapNumber int) (*models.Lap, error) {
	var lap models.Lap
	query := `
		SELECT l.lap_number, l.session_id, l.lap_time
		FROM laps l
		JOIN sessions se ON l.session_id=se.session_id
		JOIN stages s ON se.stage_id=s.stage_id
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id=t.team_id
		WHERE l.session_id=$1 AND l.lap_number=$2 AND t.user_id=$3
	`
	err := db.QueryRow(query, sessionID, lapNumber, userID).Scan(&lap.LapNumber, &lap.SessionID, &lap.LapTime)
	if err != nil {
		return nil, err
	}
	return &lap, nil
}

// UpdateLap updates a lap's time
func UpdateLap(db *sql.DB, userID, sessionID, lapNumber int, newLapTime string) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM laps l
		JOIN sessions se ON l.session_id=se.session_id
		JOIN stages s ON se.stage_id=s.stage_id
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id=t.team_id
		WHERE l.session_id=$1 AND l.lap_number=$2 AND t.user_id=$3
	`, sessionID, lapNumber, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or lap not found")
	}

	query := `UPDATE laps SET lap_time=$1 WHERE session_id=$2 AND lap_number=$3`
	_, err = db.Exec(query, newLapTime, sessionID, lapNumber)
	return err
}

// DeleteLap deletes a lap
func DeleteLap(db *sql.DB, userID, sessionID, lapNumber int) error {
	query := `
		DELETE FROM laps
		USING sessions, stages, championships, teams
		WHERE laps.session_id = sessions.session_id
		AND sessions.stage_id = stages.stage_id
		AND stages.championship_id = championships.championship_id
		AND championships.team_id = teams.team_id
		AND teams.user_id = $1
		AND laps.session_id = $2
		AND laps.lap_number = $3
	`
	_, err := db.Exec(query, userID, sessionID, lapNumber)
	return err
}
