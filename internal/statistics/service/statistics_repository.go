package service

import "github.com/babon21/statistics-counter-service/internal/statistics/domain"

type StatisticsRepository interface {
	SaveStatistics(statistic domain.Statistics) error
	GetStatisticsList(from string, to string, sortField SortField, sortOrder SortOrder) ([]domain.Statistics, error)
	ResetStatistics() error
}
