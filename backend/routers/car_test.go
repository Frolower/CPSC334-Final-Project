package routers

import (
	"Ariadne_Management/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAssignCarToTeamHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`INSERT INTO cars`).WithArgs("FOWFAO429OSBOG", "Audi", "RS 3 LMS TCR GEN 2", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	car := models.Car{
		ChassisNumber: "FOWFAO429OSBOG",
		Make:          "Audi",
		Model:         "RS 3 LMS TCR GEN 2",
	}
	body, _ := json.Marshal(car)

	req, _ := http.NewRequest(http.MethodPost, "/teams/1/cars", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/teams/:team_id/cars", AssignCarToTeamHandler(db))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Car assigned successfully")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCarHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`DELETE FROM cars`).WithArgs("FOWFAO429OSBOG").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, _ := http.NewRequest(http.MethodDelete, "/cars/FOWFAO429OSBOG", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/cars/:chassis_number", DeleteCarHandler(db))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Car deleted successfully")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCarsByTeamHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"chassis_number", "make", "model", "team_id"}).
		AddRow("FOWFAO429OSBOG", "Audi", "RS 3 LMS TCR GEN 2", 1).
		AddRow("FPAOEV4240OJN3", "Porsche", "911 992 GT3 RS", 1)

	mock.ExpectQuery(`SELECT chassis_number, make, model, team_id FROM cars WHERE team_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	req, _ := http.NewRequest(http.MethodGet, "/teams/1/cars", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/teams/:team_id/cars", GetCarsByTeamHandler(db))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "FOWFAO429OSBOG")
	assert.Contains(t, w.Body.String(), "FPAOEV4240OJN3")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCarHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`UPDATE cars SET`).WithArgs("Audi", "RS 3 LMS TCR GEN 2", 1, "FOWFAO429OSBOG").
		WillReturnResult(sqlmock.NewResult(1, 1))

	car := models.Car{
		Make:   "Audi",
		Model:  "RS 3 LMS TCR GEN 2",
		TeamID: 1,
	}
	body, _ := json.Marshal(car)

	req, _ := http.NewRequest(http.MethodPut, "/cars/FOWFAO429OSBOG", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r := gin.Default()
	r.PUT("/cars/:chassis_number", UpdateCarHandler(db))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Car updated successfully")
	assert.NoError(t, mock.ExpectationsWereMet())
}
