package services

import (
	"Ariadne_Management/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAssignCarToTeam(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`INSERT INTO cars`).WithArgs("FOWFAO429OSBOG", "Audi", "RS3 LMS TCR GEN 2", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	car := &models.Car{
		ChassisNumber: "FOWFAO429OSBOG",
		Make:          "Audi",
		Model:         "RS3 LMS TCR GEN 2",
	}

	err := AssignCarToTeam(db, 1, car)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCar(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`DELETE FROM cars`).WithArgs("FOWFAO429OSBOG").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteCar(db, "FOWFAO429OSBOG")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCarsByTeamID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"chassis_number", "make", "model", "team_id"}).
		AddRow("FOWFAO429OSBOG", "Audi", "RS3 LMS TCR GEN 2", 1).
		AddRow("FPAOEV4240OJN3", "Porshce", "911 992 GT3 RS", 1)

	mock.ExpectQuery(`SELECT chassis_number, make, model, team_id FROM cars WHERE team_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	cars, err := GetCarsByTeamID(db, 1)
	assert.NoError(t, err)
	assert.Len(t, cars, 2)
	assert.Equal(t, "FOWFAO429OSBOG", cars[0].ChassisNumber)
	assert.Equal(t, "FPAOEV4240OJN3", cars[1].ChassisNumber)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCar(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`UPDATE cars SET`).WithArgs("Audi", "RS3 LMS TCR GEN 2", 1, "FOWFAO429OSBOG").
		WillReturnResult(sqlmock.NewResult(1, 1))

	car := &models.Car{
		Make:   "Audi",
		Model:  "RS3 LMS TCR GEN 2",
		TeamID: 1,
	}

	err := UpdateCar(db, "FOWFAO429OSBOG", car)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
