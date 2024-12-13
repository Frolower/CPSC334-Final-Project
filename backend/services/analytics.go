// services/analytics.go (example)
package services

import (
	"database/sql"
)

func GetAveragePilotFinishPosition(db *sql.DB, documentNumber int, firstName, lastName string) (float64, error) {
	query := `
        SELECT AVG(result::float) 
        FROM pilotInChampionship 
        WHERE document_number=$1 AND first_name=$2 AND last_name=$3
    `
	var avg float64
	err := db.QueryRow(query, documentNumber, firstName, lastName).Scan(&avg)
	if err != nil {
		return 0, err
	}
	return avg, nil
}

// Implement similarly for other analytics queries.
