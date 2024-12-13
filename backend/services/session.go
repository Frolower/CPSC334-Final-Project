package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
)

// CreateSession creates a session linked to a stage
func CreateSession(db *sql.DB, userID int, session *models.Session) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM stages s
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE s.stage_id=$1 AND t.user_id=$2
	`, session.StageID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or stage not found")
	}

	query := `INSERT INTO sessions (stage_id, type, session_date, start_time, weather, temperature, humidity)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING session_id`
	return db.QueryRow(query, session.StageID, session.Type, session.SessionDate, session.StartTime, session.Weather, session.Temperature, session.Humidity).Scan(&session.SessionID)
}

func GetSessionsByUser(db *sql.DB, userID int) ([]models.Session, error) {
	var sessions []models.Session
	query := `
		SELECT se.session_id, se.stage_id, se.type, se.session_date, se.start_time, se.weather, se.temperature, se.humidity
		FROM sessions se
		JOIN stages s ON se.stage_id = s.stage_id
		JOIN championships ch ON s.championship_id = ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE t.user_id = $1
	`
	rows, err := db.Query(query, userID)
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

func GetSessionByID(db *sql.DB, userID, sessionID int) (*models.Session, error) {
	var sess models.Session
	query := `
		SELECT se.session_id, se.stage_id, se.type, se.session_date, se.start_time, se.weather, se.temperature, se.humidity
		FROM sessions se
		JOIN stages s ON se.stage_id = s.stage_id
		JOIN championships ch ON s.championship_id = ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE se.session_id=$1 AND t.user_id=$2
	`
	err := db.QueryRow(query, sessionID, userID).Scan(&sess.SessionID, &sess.StageID, &sess.Type, &sess.SessionDate, &sess.StartTime, &sess.Weather, &sess.Temperature, &sess.Humidity)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func UpdateSession(db *sql.DB, userID, sessionID int, session *models.Session) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM sessions se
		JOIN stages s ON se.stage_id=s.stage_id
		JOIN championships ch ON s.championship_id=ch.championship_id
		JOIN teams t ON ch.team_id=t.team_id
		WHERE se.session_id=$1 AND t.user_id=$2
	`, sessionID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or session not found")
	}

	query := `UPDATE sessions SET type=$1, session_date=$2, start_time=$3, weather=$4, temperature=$5, humidity=$6 WHERE session_id=$7`
	_, err = db.Exec(query, session.Type, session.SessionDate, session.StartTime, session.Weather, session.Temperature, session.Humidity, sessionID)
	return err
}

func DeleteSession(db *sql.DB, userID, sessionID int) error {
	query := `
		DELETE FROM sessions
		USING stages, championships, teams
		WHERE sessions.stage_id = stages.stage_id
		AND stages.championship_id = championships.championship_id
		AND championships.team_id = teams.team_id
		AND teams.user_id = $1
		AND sessions.session_id = $2
	`
	_, err := db.Exec(query, userID, sessionID)
	return err
}
