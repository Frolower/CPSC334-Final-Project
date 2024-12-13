package services

import (
	"Ariadne_Management/models"
	"database/sql"
)

// CreateSession creates a new session
func CreateSession(db *sql.DB, session *models.Session) error {
	query := `INSERT INTO sessions (stage_id, type, session_date, start_time, weather, temperature, humidity)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING session_id`
	return db.QueryRow(query, session.StageID, session.Type, session.SessionDate, session.StartTime, session.Weather, session.Temperature, session.Humidity).Scan(&session.SessionID)
}

// GetSessions retrieves all sessions
func GetSessions(db *sql.DB) ([]models.Session, error) {
	var sessions []models.Session
	query := `
		SELECT session_id, stage_id, type, session_date, start_time, weather, temperature, humidity
		FROM sessions
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sess models.Session
		if err := rows.Scan(&sess.SessionID, &sess.StageID, &sess.Type, &sess.SessionDate, &sess.StartTime, &sess.Weather, &sess.Temperature, &sess.Humidity); err != nil {
			return nil, err
		}
		sessions = append(sessions, sess)
	}
	return sessions, nil
}

// GetSessionsByStageID retrieves all sessions for a given stage_id
func GetSessionsByStageID(db *sql.DB, stageID int) ([]models.Session, error) {
	var sessions []models.Session
	query := `
		SELECT session_id, stage_id, type, session_date, start_time, weather, temperature, humidity
		FROM sessions
		WHERE stage_id = $1
	`
	rows, err := db.Query(query, stageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sess models.Session
		if err := rows.Scan(&sess.SessionID, &sess.StageID, &sess.Type, &sess.SessionDate, &sess.StartTime, &sess.Weather, &sess.Temperature, &sess.Humidity); err != nil {
			return nil, err
		}
		sessions = append(sessions, sess)
	}
	return sessions, nil
}

// UpdateSession updates a session
func UpdateSession(db *sql.DB, sessionID int, session *models.Session) error {
	query := `UPDATE sessions SET type=$1, session_date=$2, start_time=$3, weather=$4, temperature=$5, humidity=$6 WHERE session_id=$7`
	_, err := db.Exec(query, session.Type, session.SessionDate, session.StartTime, session.Weather, session.Temperature, session.Humidity, sessionID)
	return err
}

// DeleteSession deletes a session
func DeleteSession(db *sql.DB, sessionID int) error {
	query := `DELETE FROM sessions WHERE session_id=$1`
	_, err := db.Exec(query, sessionID)
	return err
}
