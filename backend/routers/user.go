package routers

import (
	"Ariadne_Management/models"
	servicies "Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// RegisterUser handles the user registration logic
func RegisterUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Delegate the user registration logic to the services layer
		err := servicies.CreateUser(db, &user)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		// Fetch the newly created user ID
		userID, err := servicies.GetUserIDByUsername(db, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user ID"})
			return
		}

		// Assign the ID to the user object
		user.ID = userID

		// Generate a JWT token for the newly created user
		token, err := servicies.GenerateJWT(&user)
		if err != nil {
			log.Printf("Error generating JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
			return
		}

		// Return the JWT token along with a success message
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"token":   token,
		})
	}
}

// LoginUser handles user login, verifying credentials and generating a JWT token
func LoginUserHandler(db *sql.DB) gin.HandlerFunc {
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
			// If the login failed, return an unauthorized error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Fetch the user ID after successful login (we are guaranteed to find the user by username)
		userID, err := servicies.GetUserIDByUsername(db, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user ID"})
			return
		}
		user.ID = userID // Assign the fetched ID to the user struct

		// If status is true, generate a JWT token
		token, err := servicies.GenerateJWT(&user)
		if err != nil {
			log.Printf("Error generating JWT: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
			return
		}

		// Return the JWT token in the response
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully",
			"token":   token,
		})
	}
}
