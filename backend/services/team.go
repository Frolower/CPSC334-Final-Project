package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

func CreateTeam(db *sql.DB, team *models.Team) error {
	query := `INSERT INTO teams (team_id, user_id, team_name) VALUES ($1, $2, $3)`
	result, err := db.Exec(query, team.Team_ID, team.User_ID, team.Team_name)
	log.Println(result)
	return err
}
