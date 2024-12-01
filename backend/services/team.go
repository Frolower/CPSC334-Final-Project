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

// GetTeamsByUserID retrieves all teams associated with a user
func GetTeamsByUserID(db *sql.DB, userID int) ([]models.Team, error) {
	var teams []models.Team

	// Query to get all teams for the user
	query := `SELECT team_id, user_id, team_name FROM teams WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows and map each row to a Team struct
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
