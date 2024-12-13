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

// POST /assignCarToTeam/:team_id
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

// DELETE /deleteCar/:chassis_number
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

// GET /getCarsByUser
// Since we're removing ownership checks, this now returns all cars
func GetCarsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cars, err := services.GetCars(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cars"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	}
}
