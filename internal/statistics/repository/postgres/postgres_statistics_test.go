package postgres_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"github.com/babon21/statistics-counter-service/internal/statistics/repository/postgres"
	"github.com/babon21/statistics-counter-service/internal/statistics/service"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetStatisticsList_Desc(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockStatistics := []domain.Statistics{
		domain.NewStatistics(domain.SimpleStatistics{
			Date:   "2002-03-30",
			Views:  1,
			Clicks: 2,
			Cost:   3,
		}),
		domain.NewStatistics(domain.SimpleStatistics{
			Date:   "2002-03-28",
			Views:  3,
			Clicks: 4,
			Cost:   5,
		}),
	}

	rows := sqlmock.NewRows([]string{"date", "views", "clicks", "cost", "cpc", "cpm"}).
		AddRow(mockStatistics[0].Date, mockStatistics[0].Views, mockStatistics[0].Clicks, mockStatistics[0].Cost, mockStatistics[0].Cpc, mockStatistics[0].Cpm).
		AddRow(mockStatistics[1].Date, mockStatistics[1].Views, mockStatistics[1].Clicks, mockStatistics[1].Cost, mockStatistics[1].Cpc, mockStatistics[1].Cpm)

	sortField := service.DateField
	sortOrder := service.DescOrder

	from := "2002-01-30"
	to := "2002-04-30"

	query := fmt.Sprintf("SELECT date,views,clicks,cost,cpc,cpm FROM statistics WHERE date BETWEEN '%s' AND '%s' ORDER BY %s %s", from, to, sortField, strings.ToUpper(string(sortOrder)))

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresStatisticsRepository(db)

	list, err := a.GetStatisticsList(from, to, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetStatisticsList_Asc(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockStatistics := []domain.Statistics{
		domain.NewStatistics(domain.SimpleStatistics{
			Date:   "2002-03-28",
			Views:  1,
			Clicks: 2,
			Cost:   3,
		}),
		domain.NewStatistics(domain.SimpleStatistics{
			Date:   "2002-03-30",
			Views:  3,
			Clicks: 4,
			Cost:   5,
		}),
	}

	rows := sqlmock.NewRows([]string{"date", "views", "clicks", "cost", "cpc", "cpm"}).
		AddRow(mockStatistics[0].Date, mockStatistics[0].Views, mockStatistics[0].Clicks, mockStatistics[0].Cost, mockStatistics[0].Cpc, mockStatistics[0].Cpm).
		AddRow(mockStatistics[1].Date, mockStatistics[1].Views, mockStatistics[1].Clicks, mockStatistics[1].Cost, mockStatistics[1].Cpc, mockStatistics[1].Cpm)

	sortField := service.DateField
	sortOrder := service.AscOrder

	from := "2002-01-30"
	to := "2002-04-30"

	query := fmt.Sprintf("SELECT date,views,clicks,cost,cpc,cpm FROM statistics WHERE date BETWEEN '%s' AND '%s' ORDER BY %s %s", from, to, sortField, strings.ToUpper(string(sortOrder)))

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresStatisticsRepository(db)

	list, err := a.GetStatisticsList(from, to, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestSaveStatistics(t *testing.T) {
	s := domain.NewStatistics(domain.SimpleStatistics{
		Date:   "2002-03-28",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	})

	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mock.ExpectExec("INSERT INTO statistics").WithArgs(s.Date, s.Views, s.Clicks, s.Cost, s.Cpc, s.Cpm).WillReturnResult(sqlmock.NewResult(2, 1))

	a := postgres.NewPostgresStatisticsRepository(db)

	err = a.SaveStatistics(s)
	assert.NoError(t, err)
}

func TestResetStatistics(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mock.ExpectExec("DELETE FROM statistics").WillReturnResult(sqlmock.NewResult(2, 1))

	a := postgres.NewPostgresStatisticsRepository(db)

	err = a.ResetStatistics()
	assert.NoError(t, err)
}
