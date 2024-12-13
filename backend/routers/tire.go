package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AssignTireToCarHandler
func AssignTireToCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")

		var tire models.Tire
		if err := c.ShouldBindJSON(&tire); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		err := services.AssignTireToCar(db, chassisNumber, &tire)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign tire to car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tire assigned successfully"})
	}
}

// DeleteTireHandler
func DeleteTireHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tireID := c.Param("tire_id")
		err := services.DeleteTire(db, tireID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete tire"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tire deleted successfully"})
	}
}

// GetTiresByCarHandler
func GetTiresByCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		tires, err := services.GetTiresByChassisNumber(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tires"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tires": tires})
	}
}

// UpdateTireHandler
func UpdateTireHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tireID := c.Param("tire_id")

		var tire models.Tire
		if err := c.ShouldBindJSON(&tire); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		err := services.UpdateTire(db, tireID, &tire)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update tire"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tire updated successfully"})
	}
}
