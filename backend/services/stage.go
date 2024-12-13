package services

import (
	"Ariadne_Management/models"
	"database/sql"
)

// CreateStage creates a new stage
func CreateStage(db *sql.DB, stage *models.Stage) error {
	query := `INSERT INTO stages (stage_number, championship_id, track, start_date, end_date) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING stage_id`
	return db.QueryRow(query, stage.StageNumber, stage.ChampionshipID, stage.Track, stage.StartDate, stage.EndDate).Scan(&stage.StageID)
}

// GetStages retrieves all stages (no user filtering)
func GetStages(db *sql.DB) ([]models.Stage, error) {
	var stages []models.Stage
	query := `
		SELECT stage_id, stage_number, championship_id, track, start_date, end_date
		FROM stages
	`
	rows, err := db.Query(query)
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

// GetStagesByChampionshipID retrieves all stages for a given championship_id
func GetStagesByChampionshipID(db *sql.DB, championshipID int) ([]models.Stage, error) {
	var stages []models.Stage
	query := `
		SELECT stage_id, stage_number, championship_id, track, start_date, end_date
		FROM stages
		WHERE championship_id = $1
	`
	rows, err := db.Query(query, championshipID)
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

// UpdateStage updates a stage
func UpdateStage(db *sql.DB, stageID int, stage *models.Stage) error {
	query := `UPDATE stages SET stage_number=$1, track=$2, start_date=$3, end_date=$4 WHERE stage_id=$5`
	_, err := db.Exec(query, stage.StageNumber, stage.Track, stage.StartDate, stage.EndDate, stageID)
	return err
}

// DeleteStage deletes a stage
func DeleteStage(db *sql.DB, stageID int) error {
	query := `DELETE FROM stages WHERE stage_id=$1`
	_, err := db.Exec(query, stageID)
	return err
}
