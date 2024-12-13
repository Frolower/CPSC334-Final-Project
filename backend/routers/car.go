package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// AssignCarToTeamHandler
func AssignCarToTeamHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		teamIDStr := c.Param("team_id")
		teamID, err := strconv.Atoi(teamIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
			return
		}

		var car models.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err = services.AssignCarToTeam(db, teamID, &car)
		if err != nil {
			log.Printf("Error assigning car to team: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign car to team"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Car assigned successfully"})
	}
}

// DeleteCarHandler
func DeleteCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		err := services.DeleteCar(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete car"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
	}
}

// GetCarsByTeamHandler
func GetCarsByTeamHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		teamIDStr := c.Param("team_id")
		teamID, err := strconv.Atoi(teamIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
			return
		}

		cars, err := services.GetCarsByTeamID(db, teamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cars"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	}
}

// UpdateCarHandler
func UpdateCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")

		var car models.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := services.UpdateCar(db, chassisNumber, &car)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
	}
}
