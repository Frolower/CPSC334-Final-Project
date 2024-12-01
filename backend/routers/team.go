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

func CreateTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var team models.Team
		if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Delegate the user registration logic to the services layer
		if err := servicies.CreateTeam(db, &team); err != nil {
			log.Printf("Error creating team: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create team"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
	}
}

// GetUserTeams handles the GET request to fetch teams for the logged-in user
func GetUserTeams(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the request header
		tokenString := c.GetHeader("Authorization")
		userID, err := servicies.ExtractUserIDFromToken(tokenString) // extractUserIDFromToken is a function to decode the JWT
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Fetch teams for the user from the database
		teams, err := servicies.GetTeamsByUserID(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching teams"})
			return
		}

		// Return the list of teams
		c.JSON(http.StatusOK, gin.H{
			"teams": teams,
		})
	}
}
