// routers/analytics.go (example for an analytics endpoint)
package routers

import (
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAvgPilotFinishPositionHandler
func GetAvgPilotFinishPositionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		documentNumberStr := c.Param("document_number")
		firstName := c.Param("first_name")
		lastName := c.Param("last_name")

		documentNumber, err := strconv.Atoi(documentNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document_number"})
			return
		}

		avg, err := services.GetAveragePilotFinishPosition(db, documentNumber, firstName, lastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not compute average"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"average_finish_position": avg})
	}
}
