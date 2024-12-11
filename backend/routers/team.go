package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateTeam handles creating a team for an authenticated user
func CreateTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// Extract the user ID from the JWT token
		userID, err := services.ExtractUserIDFromToken(tokenString)
		if err != nil {
			log.Printf("Error extracting user ID from token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Bind the incoming request JSON to the Team struct
		var team models.Team
		if err := c.ShouldBindJSON(&team); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Ensure team_name is provided
		if team.Team_name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Team name is required"})
			return
		}

		// Log the team data to check the incoming request
		log.Printf("Creating team with name: %s for user ID: %d", team.Team_name, userID)

		// Assign the user ID to the team struct
		team.User_ID = userID

		// Call the service layer to insert the team into the database
		if err := services.CreateTeam(db, &team); err != nil {
			log.Printf("Error creating team: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create team"})
			return
		}

		// Return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
	}
}
