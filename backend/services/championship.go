package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"fmt"
)

// CreateChampionship creates a new championship linked to a user's team
func CreateChampionship(db *sql.DB, userID int, champ *models.Championship) error {
	// Verify that the team belongs to the user
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM teams WHERE team_id=$1 AND user_id=$2", champ.TeamID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or team not found")
	}

	query := `INSERT INTO championships (team_id, championship_name, team_standings) VALUES ($1, $2, $3) RETURNING championship_id`
	return db.QueryRow(query, champ.TeamID, champ.ChampionshipName, champ.TeamStandings).Scan(&champ.ChampionshipID)
}

// GetChampionshipsByUser retrieves all championships belonging to the user
func GetChampionshipsByUser(db *sql.DB, userID int) ([]models.Championship, error) {
	var champs []models.Championship
	query := `
		SELECT ch.championship_id, ch.team_id, ch.championship_name, ch.team_standings
		FROM championships ch
		JOIN teams t ON ch.team_id = t.team_id
		WHERE t.user_id = $1
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Championship
		if err := rows.Scan(&c.ChampionshipID, &c.TeamID, &c.ChampionshipName, &c.TeamStandings); err != nil {
			return nil, err
		}
		champs = append(champs, c)
	}
	return champs, nil
}

// GetChampionshipByID retrieves a single championship by ID, ensuring user ownership
func GetChampionshipByID(db *sql.DB, userID, championshipID int) (*models.Championship, error) {
	var c models.Championship
	query := `
		SELECT ch.championship_id, ch.team_id, ch.championship_name, ch.team_standings
		FROM championships ch
		JOIN teams t ON ch.team_id = t.team_id
		WHERE ch.championship_id = $1 AND t.user_id = $2
	`
	err := db.QueryRow(query, championshipID, userID).Scan(&c.ChampionshipID, &c.TeamID, &c.ChampionshipName, &c.TeamStandings)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// UpdateChampionship updates a championship
func UpdateChampionship(db *sql.DB, userID, championshipID int, champ *models.Championship) error {
	// Ensure ownership
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM championships ch
		JOIN teams t ON ch.team_id = t.team_id
		WHERE ch.championship_id=$1 AND t.user_id=$2
	`, championshipID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("unauthorized or championship not found")
	}

	query := `UPDATE championships SET championship_name=$1, team_standings=$2 WHERE championship_id=$3`
	_, err = db.Exec(query, champ.ChampionshipName, champ.TeamStandings, championshipID)
	return err
}

// DeleteChampionship deletes a championship
func DeleteChampionship(db *sql.DB, userID, championshipID int) error {
	query := `
		DELETE FROM championships 
		USING teams 
		WHERE championships.team_id = teams.team_id
		AND teams.user_id = $1
		AND championships.championship_id = $2
	`
	_, err := db.Exec(query, userID, championshipID)
	return err
}
