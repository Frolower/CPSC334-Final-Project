// services/team.go
package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"log"
)

func CreateTeam(db *sql.DB, team *models.Team) error {
	query := `INSERT INTO teams (user_id, team_name) VALUES ($1, $2)`
	_, err := db.Exec(query, team.User_ID, team.Team_name)
	if err != nil {
		log.Printf("Error inserting team into database: %v", err)
		return err
	}
	return nil
}

func GetTeamsByUserID(db *sql.DB, userID int) ([]models.Team, error) {
	var teams []models.Team
	query := `SELECT team_id, user_id, team_name FROM teams WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.Team_ID, &team.User_ID, &team.Team_name)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}

func DeleteTeam(db *sql.DB, userID, teamID int) error {
	// Delete only if the team belongs to the user
	_, err := db.Exec(`DELETE FROM teams WHERE team_id=$1 AND user_id=$2`, teamID, userID)
	return err
}
