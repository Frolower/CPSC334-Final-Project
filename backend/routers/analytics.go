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

		lapNumber, lapTime, lapSeconds, err := services.GetFastestLap(db, sessionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve fastest lap"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"lap_number":       lapNumber,
			"lap_time":         lapTime,
			"lap_time_seconds": lapSeconds,
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

// GetTireCountForCarHandler
func GetTireCountForCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		count, err := services.GetTireCountForCar(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tire count"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tire_count": count})
	}
}

// GetAverageTreadForCarHandler
func GetAverageTreadForCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")
		avg, err := services.GetAverageTreadForCar(db, chassisNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve average tread"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"average_tread_remaining": avg})
	}
}

// GetTiresCountByTreadAndCompoundAndCarHandler
func GetTiresCountByTreadAndCompoundAndCarHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		chassisNumber := c.Param("chassis_number")

		treadStr := c.Param("tread")
		tread, err := strconv.ParseFloat(treadStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tread value"})
			return
		}

		compound := c.Param("compound")

		count, err := services.GetTiresCountByTreadAndCompoundAndCar(db, chassisNumber, tread, compound)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tires count by tread, compound, and car"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tires_count": count})
	}
}
