package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Create
func CreateSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		if err := services.CreateSession(db, &sess); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session created", "session_id": sess.SessionID})
	}
}

// GET all sessions
func GetSessionsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions, err := services.GetSessions(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching sessions"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sessions": sessions})
	}
}

// GET
func GetSessionByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		sess, err := services.GetSessionByID(db, sessionID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"session": sess})
	}
}

// Update
func UpdateSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		if err := services.UpdateSession(db, sessionID, &sess); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session updated"})
	}
}

// DELETE
func DeleteSessionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		if err := services.DeleteSession(db, sessionID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
	}
}
