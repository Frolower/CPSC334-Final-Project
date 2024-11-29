// /config/db.go

package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
)

func InitDB() (*sql.DB, error) {
	// Get the database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
		return nil, err
	}

	log.Println("Connected to the database successfully.")
	return db, nil
}
