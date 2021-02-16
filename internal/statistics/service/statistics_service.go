package service

import (
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"time"
)

type StatisticsService interface {
	SaveStatistics(statistic domain.Statistics) error
	GetStatisticsList(from string, to string, sortField SortField, sortOrder SortOrder) ([]domain.Statistics, error)
	ResetStatistics() error
}

type statisticsService struct {
	statisticRepo StatisticsRepository
}

func NewStatisticsService(repository StatisticsRepository) StatisticsService {
	return &statisticsService{statisticRepo: repository}
}

func (s *statisticsService) SaveStatistics(statistics domain.Statistics) error {
	return s.statisticRepo.SaveStatistics(statistics)
}

func (s *statisticsService) GetStatisticsList(from string, to string, sortField SortField, sortOrder SortOrder) ([]domain.Statistics, error) {
	list, err := s.statisticRepo.GetStatisticsList(from, to, sortField, sortOrder)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(list); i++ {
		t, _ := time.Parse(time.RFC3339, list[i].Date)
		list[i].Date = t.Format(LayoutISO)
	}

	return list, nil
}

func (s *statisticsService) ResetStatistics() error {
	return s.statisticRepo.ResetStatistics()
}
