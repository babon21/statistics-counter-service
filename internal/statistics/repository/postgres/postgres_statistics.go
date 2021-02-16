package postgres

import (
	"fmt"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"github.com/babon21/statistics-counter-service/internal/statistics/service"
	"github.com/jmoiron/sqlx"
)

type postgresStatisticsRepository struct {
	Conn *sqlx.DB
}

func (p *postgresStatisticsRepository) SaveStatistics(s domain.Statistics) error {
	_, err := p.Conn.Exec("INSERT INTO statistics(date, views, clicks, cost, cpc, cpm) VALUES ($1,$2,$3,$4,$5,$6)", s.Date, s.Views, s.Clicks, s.Cost, s.Cpc, s.Cpm)
	return err
}

func (p *postgresStatisticsRepository) GetStatisticsList(from string, to string, sortField service.SortField, sortOrder service.SortOrder) ([]domain.Statistics, error) {
	getListQuery := formGetListQuery(from, to, sortField, sortOrder)
	statisticsList := make([]domain.Statistics, 0, 1)
	err := p.Conn.Select(&statisticsList, getListQuery)
	return statisticsList, err
}

func formGetListQuery(from string, to string, sortField service.SortField, sortOrder service.SortOrder) string {
	var order string
	switch sortOrder {
	case service.AscOrder:
		order = "ASC"
	case service.DescOrder:
		order = "DESC"
	}

	return fmt.Sprintf("SELECT date,views,clicks,cost,cpc,cpm FROM statistics WHERE date BETWEEN '%s' AND '%s' ORDER BY %s %s", from, to, sortField, order)
}

func (p *postgresStatisticsRepository) ResetStatistics() error {
	_, err := p.Conn.Exec("DELETE FROM statistics")
	return err
}

func NewPostgresStatisticsRepository(conn *sqlx.DB) service.StatisticsRepository {
	return &postgresStatisticsRepository{conn}
}
