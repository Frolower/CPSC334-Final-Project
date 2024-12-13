package services

import (
	"Ariadne_Management/models"
	"database/sql"
)

// CreateChampionship creates a new championship
func CreateChampionship(db *sql.DB, champ *models.Championship) error {
	query := `INSERT INTO championships (team_id, championship_name, team_standings)
	          VALUES ($1, $2, $3) RETURNING championship_id`
	return db.QueryRow(query, champ.TeamID, champ.ChampionshipName, champ.TeamStandings).Scan(&champ.ChampionshipID)
}

// GetChampionships retrieves all championships (no user filtering)
func GetChampionships(db *sql.DB) ([]models.Championship, error) {
	var champs []models.Championship
	query := `SELECT championship_id, team_id, championship_name, team_standings FROM championships`
	rows, err := db.Query(query)
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

// GetChampionshipByID retrieves a single championship by ID
func GetChampionshipByID(db *sql.DB, championshipID int) (*models.Championship, error) {
	var c models.Championship
	query := `SELECT championship_id, team_id, championship_name, team_standings
	          FROM championships WHERE championship_id=$1`
	err := db.QueryRow(query, championshipID).Scan(&c.ChampionshipID, &c.TeamID, &c.ChampionshipName, &c.TeamStandings)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// UpdateChampionship updates a championship
func UpdateChampionship(db *sql.DB, championshipID int, champ *models.Championship) error {
	query := `UPDATE championships SET championship_name=$1, team_standings=$2 WHERE championship_id=$3`
	_, err := db.Exec(query, champ.ChampionshipName, champ.TeamStandings, championshipID)
	return err
}

// DeleteChampionship deletes a championship
func DeleteChampionship(db *sql.DB, championshipID int) error {
	query := `DELETE FROM championships WHERE championship_id=$1`
	_, err := db.Exec(query, championshipID)
	return err
}
