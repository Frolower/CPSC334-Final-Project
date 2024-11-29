// /services/user.go

package services

import (
	"Ariadne_Management/models"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// CreateUser handles user registration logic
func CreateUser(db *sql.DB, user *models.User) error {
	var result sql.Result
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert user into DB
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	result, err = db.Exec(query, user.Username, user.Password)
	log.Println(result)
	return err
}
