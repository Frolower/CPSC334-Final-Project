package services

import (
	"Ariadne_Management/models"
	servicies "Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Import pq for PostgreSQL support
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
