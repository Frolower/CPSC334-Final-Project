package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"log"
)

func AssignPartToCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		chassisNumber := c.Param("chassis_number")

		var part models.Part
		if err := c.ShouldBindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := services.AssignPartToCar(db, userID.(int), chassisNumber, &part)
		if err != nil {
			log.Printf("Error assigning part to car: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign part to car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Part assigned successfully"})
	}
}

// DeletePartHandler
func DeletePartHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		partID := c.Param("part_id")
		err := services.DeletePart(db, userID.(int), partID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete part"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Part deleted successfully"})
	}
}

// GetPartsByUserHandler
func GetPartsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		parts, err := services.GetPartsByUser(db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch parts"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parts": parts})
	}
}
