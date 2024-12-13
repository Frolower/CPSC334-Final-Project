package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		stageIDStr := c.Param("stage_id")
		stageID, err := strconv.Atoi(stageIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stage ID"})
			return
		}

		var sess models.Session
		if err := c.ShouldBindJSON(&sess); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		sess.StageID = stageID

		if err := services.CreateSession(db, userID.(int), &sess); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session created", "session_id": sess.SessionID})
	}
}

func GetSessionsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessions, err := services.GetSessionsByUser(db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching sessions"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sessions": sessions})
	}
}

func GetSessionByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		sess, err := services.GetSessionByID(db, userID.(int), sessionID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"session": sess})
	}
}

func UpdateSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		var sess models.Session
		if err := c.ShouldBindJSON(&sess); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := services.UpdateSession(db, userID.(int), sessionID, &sess); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session updated"})
	}
}

func DeleteSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		if err := services.DeleteSession(db, userID.(int), sessionID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
	}
}
