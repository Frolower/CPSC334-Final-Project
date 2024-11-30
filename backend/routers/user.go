package services

import (
	"Ariadne_Management/models"
	servicies "Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Delegate the user registration logic to the services layer
		if err := servicies.CreateUser(db, &user); err != nil {
			log.Printf("Error creating user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

// Function to generate a JWT token
//func fenerateJWT(user *models.User) (string, error) {
//	// Define token expiration time (e.g., 1 hour)
//	expirationTime := time.Now().Add(1 * time.Hour)
//	// Create the JWT claims
//	claims := &jwt.StandardClaims{
//		Subject:   user.Email,         // The user email as subject
//		ExpiresAt: expirationTime.Unix(),
//		IssuedAt:  time.Now().Unix(),
//	}
//
//	// Create a new JWT token using the claims
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	// Sign the token using your secret key
//	signedToken, err := token.SignedString([]byte("yourSecretKey"))
//	if err != nil {
//		return "", err
//	}
//	return signedToken, nil
//}

// LoginUser handles user login
func LoginUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		// Parse the incoming JSON request into the user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Use the service function to verify if the user exists and the password is correct
		status, err := servicies.VerifyUserLogIn(db, &user)
		if err != nil {
			log.Printf("Error verifying user login: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log in the user"})
			return
		}

		// Check if the status is true or false
		if !status {
			// If the login failed (status is false), return an unauthorized error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// If status is true (login success), return a success message
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
	}
}
