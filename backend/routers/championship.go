package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Create a championship
func CreateChampionshipHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var champ models.Championship
		if err := c.ShouldBindJSON(&champ); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := services.CreateChampionship(db, &champ); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create championship"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Championship created", "championship_id": champ.ChampionshipID})
	}
}

// GET all championships
func GetChampionshipsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		champs, err := services.GetChampionships(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching championships"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"championships": champs})
	}
}

// GET
func GetChampionshipByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chIDStr := c.Param("id")
		chID, err := strconv.Atoi(chIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid championship ID"})
			return
		}

		champ, err := services.GetChampionshipByID(db, chID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Championship not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"championship": champ})
	}
}

// Update
func UpdateChampionshipHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chIDStr := c.Param("id")
		chID, err := strconv.Atoi(chIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid championship ID"})
			return
		}

		var champ models.Championship
		if err := c.ShouldBindJSON(&champ); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := services.UpdateChampionship(db, chID, &champ); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update championship"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Championship updated"})
	}
}

// DELETE
func DeleteChampionshipHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chIDStr := c.Param("id")
		chID, err := strconv.Atoi(chIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid championship ID"})
			return
		}

		if err := services.DeleteChampionship(db, chID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete championship"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Championship deleted"})
	}
}
