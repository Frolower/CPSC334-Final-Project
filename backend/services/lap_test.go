package services

import (
	"Ariadne_Management/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLap(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`INSERT INTO laps`).WithArgs(1, 1, "1:20.123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	lap := &models.Lap{
		LapNumber: 1,
		SessionID: 1,
		LapTime:   "1:20.123",
	}

	err := CreateLap(db, lap)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteLap(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`DELETE FROM laps`).WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteLap(db, 1, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetLapsBySessionID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"lap_number", "lap_time"}).
		AddRow(1, "1:20.123").
		AddRow(2, "1:21.456")

	mock.ExpectQuery(`SELECT lap_number, lap_time FROM laps WHERE session_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	laps, err := GetLapsBySessionID(db, 1)
	assert.NoError(t, err)
	assert.Len(t, laps, 2)

	assert.Equal(t, 1, laps[0].LapNumber)
	assert.Equal(t, "1:20.123", laps[0].LapTime)

	assert.Equal(t, 2, laps[1].LapNumber)
	assert.Equal(t, "1:21.456", laps[1].LapTime)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateLap(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`UPDATE laps SET`).WithArgs("1:23.456", 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := UpdateLap(db, 1, 1, "1:23.456")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
