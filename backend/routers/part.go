package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /assignPartToCar/:chassis_number
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
			log.Printf("Error assigning part to car: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign part to car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Part assigned successfully"})
	}
}

// DELETE /deletePart/:part_id
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

// GET /getPartsByUser
// Since we're removing ownership checks, this now returns all parts
func GetPartsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		parts, err := services.GetParts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch parts"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parts": parts})
	}
}
