package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
)

// CreateStage creates a stage linked to a championship owned by the user
func CreateStage(db *sql.DB, userID int, stage *models.Stage) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM championships ch
		JOIN teams t ON ch.team_id = t.team_id
		WHERE ch.championship_id=$1 AND t.user_id=$2
	`, stage.ChampionshipID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or championship not found")
	}

	query := `INSERT INTO stages (stage_number, championship_id, track, start_date, end_date) 
		VALUES ($1, $2, $3, $4, $5) RETURNING stage_id`
	return db.QueryRow(query, stage.StageNumber, stage.ChampionshipID, stage.Track, stage.StartDate, stage.EndDate).Scan(&stage.StageID)
}

func GetStagesByUser(db *sql.DB, userID int) ([]models.Stage, error) {
	var stages []models.Stage
	query := `
		SELECT s.stage_id, s.stage_number, s.championship_id, s.track, s.start_date, s.end_date
		FROM stages s
		JOIN championships ch ON s.championship_id = ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE t.user_id=$1
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var st models.Stage
		if err := rows.Scan(&st.StageID, &st.StageNumber, &st.ChampionshipID, &st.Track, &st.StartDate, &st.EndDate); err != nil {
			return nil, err
		}
		stages = append(stages, st)
	}
	return stages, nil
}

func GetStageByID(db *sql.DB, userID, stageID int) (*models.Stage, error) {
	var st models.Stage
	query := `
		SELECT s.stage_id, s.stage_number, s.championship_id, s.track, s.start_date, s.end_date
		FROM stages s
		JOIN championships ch ON s.championship_id = ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE s.stage_id=$1 AND t.user_id=$2
	`
	err := db.QueryRow(query, stageID, userID).Scan(&st.StageID, &st.StageNumber, &st.ChampionshipID, &st.Track, &st.StartDate, &st.EndDate)
	if err != nil {
		return nil, err
	}
	return &st, nil
}

func UpdateStage(db *sql.DB, userID, stageID int, stage *models.Stage) error {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM stages s
		JOIN championships ch ON s.championship_id = ch.championship_id
		JOIN teams t ON ch.team_id = t.team_id
		WHERE s.stage_id=$1 AND t.user_id=$2
	`, stageID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or stage not found")
	}

	query := `UPDATE stages SET stage_number=$1, track=$2, start_date=$3, end_date=$4 WHERE stage_id=$5`
	_, err = db.Exec(query, stage.StageNumber, stage.Track, stage.StartDate, stage.EndDate, stageID)
	return err
}

func DeleteStage(db *sql.DB, userID, stageID int) error {
	query := `
		DELETE FROM stages
		USING championships, teams
		WHERE stages.championship_id = championships.championship_id
		AND championships.team_id = teams.team_id
		AND teams.user_id = $1
		AND stages.stage_id = $2
	`
	_, err := db.Exec(query, userID, stageID)
	return err
}
