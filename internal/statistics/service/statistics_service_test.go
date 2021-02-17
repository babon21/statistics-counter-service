package service_test

import (
	"errors"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain/mocks"
	"github.com/babon21/statistics-counter-service/internal/statistics/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetList(t *testing.T) {
	mockStatisticsRepo := new(mocks.StatisticsRepository)
	mockStatistics := domain.NewStatistics(domain.SimpleStatistics{
		Date:   "2002-03-30",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	})

	mockListStatistics := make([]domain.Statistics, 0)
	mockListStatistics = append(mockListStatistics, mockStatistics)

	t.Run("success", func(t *testing.T) {
		mockStatisticsRepo.On("GetStatisticsList", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(mockListStatistics, nil).Once()
		s := service.NewStatisticsService(mockStatisticsRepo)
		from := "2002-01-30"
		to := "2002-04-30"
		list, err := s.GetStatisticsList(from, to, service.DateField, service.AscOrder)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListStatistics))

		mockStatisticsRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockStatisticsRepo.On("GetStatisticsList", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, mock.Anything).
			Return(nil, errors.New("Unexpected Error")).Once()
		s := service.NewStatisticsService(mockStatisticsRepo)
		from := "2002-01-30"
		to := "2002-04-30"
		list, err := s.GetStatisticsList(from, to, service.DateField, service.AscOrder)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockStatisticsRepo.AssertExpectations(t)
	})
}

func TestSaveStatistics(t *testing.T) {
	mockStatisticsRepo := new(mocks.StatisticsRepository)
	mockStatistics := domain.NewStatistics(domain.SimpleStatistics{
		Date:   "2002-03-30",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	})

	t.Run("success", func(t *testing.T) {
		mockStatisticsRepo.On("SaveStatistics", mock.Anything).Return(nil)
		s := service.NewStatisticsService(mockStatisticsRepo)
		err := s.SaveStatistics(mockStatistics)

		assert.NoError(t, err)
		mockStatisticsRepo.AssertExpectations(t)
	})
}

func TestResetStatistics(t *testing.T) {
	mockStatisticsRepo := new(mocks.StatisticsRepository)

	t.Run("success", func(t *testing.T) {
		//mockStatisticsRepo.On("CheckExistence", mock.AnythingOfType("string")).Return(true).Once()
		mockStatisticsRepo.On("ResetStatistics").Return(nil).Once()

		s := service.NewStatisticsService(mockStatisticsRepo)
		err := s.ResetStatistics()

		assert.NoError(t, err)
		mockStatisticsRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockStatisticsRepo.On("ResetStatistics").Return(errors.New("Unexpected Error")).Once()

		s := service.NewStatisticsService(mockStatisticsRepo)
		err := s.ResetStatistics()

		assert.Error(t, err)
		mockStatisticsRepo.AssertExpectations(t)
	})
}
