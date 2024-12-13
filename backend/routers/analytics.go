package routers

import (
	services "Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetFastestLapHandler retrieves the fastest lap for a given session
func GetFastestLapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			return
		}

		lapNumber, lapTime, err := services.GetFastestLap(db, sessionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve fastest lap"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"lap_number": lapNumber,
			"lap_time":   lapTime,
		})
	}
}

// GetAverageLapHandler retrieves the average lap time for a given session
func GetAverageLapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			return
		}

		avgSeconds, err := services.GetAverageLapTime(db, sessionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not compute average lap time"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"average_lap_time_seconds": avgSeconds})
	}
}

// GetPartsCountForCarHandler retrieves the number of parts assigned to a given car
func GetPartsCountForCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		count, err := services.GetPartsCountForCar(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve parts count"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parts_count": count})
	}
}
