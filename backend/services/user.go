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
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	result, err = db.Exec(query, user.Username, user.Email, user.Password)
	log.Println(result)
	return err
}

func VerifyUserLogIn(db *sql.DB, user *models.User) (bool, error) {
	var hashedPassword string

	query := `SELECT password FROM users WHERE username = $1`
	result, err := db.Query(query, user.Username)
	if err != nil {
		log.Println(err)
		return false, err
	}

	result.Next()
	err = result.Scan(&hashedPassword) // Transfer password from *Rows into a string

	if err != nil {
		log.Println(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)) != nil {
		log.Println("Passwords do not match")
		log.Println(err)
		return false, err
	}
	return true, err
}
