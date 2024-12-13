package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AssignPartToCarHandler
func AssignPartToCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")

		var part models.Part
		if err := c.ShouldBindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := services.AssignPartToCar(db, chassisNumber, &part)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign part to car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Part assigned successfully"})
	}
}

// DeletePartHandler
func DeletePartHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		partID := c.Param("part_id")
		err := services.DeletePart(db, partID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete part"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Part deleted successfully"})
	}
}

// GetPartsByCarHandler
func GetPartsByCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		parts, err := services.GetPartsByChassisNumber(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch parts"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parts": parts})
	}
}

// UpdatePartHandler
func UpdatePartHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		partID := c.Param("part_id")

		var part models.Part
		if err := c.ShouldBindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := services.UpdatePart(db, partID, &part)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update part"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Part updated successfully"})
	}
}
