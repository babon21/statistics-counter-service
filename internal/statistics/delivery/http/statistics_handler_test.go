package http_test

import (
	"encoding/json"
	"errors"
	statisticsHttp "github.com/babon21/statistics-counter-service/internal/statistics/delivery/http"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain/mocks"
	"github.com/babon21/statistics-counter-service/pkg/delivery/http/api"
	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetStatisticsList_Ok(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	mockService.On("GetStatisticsList", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(mockAds, nil)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetStatisticsList_WithoutMediaType(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestGetStatisticsList_EmptyReqBody(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetStatisticsList_EmptySortOrderParam(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetStatisticsList_EmptySortFieldParam(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetStatisticsList_WrongSortOrderParam(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		SortOrder: "wrong",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetStatisticsList_WrongSortFieldParam(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "wrong",
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetStatisticsList_ServiceError(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	mockService.On("GetStatisticsList", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(nil, errors.New("Service error"))

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetStatisticsList_ToDateInvalid(t *testing.T) {
	var mockAd domain.Statistics
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.StatisticsService)
	mockAds := make([]domain.Statistics, 0)
	mockAds = append(mockAds, mockAd)

	j, err := json.Marshal(api.GetStatisticsRequest{
		SortField: "date",
		SortOrder: "asc",
		From:      "2020-03-13",
		To:        "2020-03-15f",
	})

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/statistics", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.GetStatisticsList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSaveStatistics_Ok(t *testing.T) {
	mockStatistics := domain.SimpleStatistics{
		Date:   "2002-03-30",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	}

	tempMockStatistics := mockStatistics
	mockService := new(mocks.StatisticsService)

	j, err := json.Marshal(tempMockStatistics)
	assert.NoError(t, err)

	mockService.On("SaveStatistics", mock.AnythingOfType("domain.Statistics")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.SaveStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestSaveStatistics_InvalidDateParam(t *testing.T) {
	mockStatistics := domain.SimpleStatistics{
		Date:   "2002-03-30f",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	}

	tempMockStatistics := mockStatistics
	mockService := new(mocks.StatisticsService)

	j, err := json.Marshal(tempMockStatistics)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.SaveStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSaveStatistics_InvalidReqBody(t *testing.T) {
	mockStatistics := domain.SimpleStatistics{
		Date:   "2002-03-30",
		Views:  -1,
		Clicks: 2,
		Cost:   3,
	}

	tempMockStatistics := mockStatistics
	mockService := new(mocks.StatisticsService)

	j, err := json.Marshal(tempMockStatistics)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.SaveStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSaveStatistics_WithoutMediaType(t *testing.T) {
	mockStatistics := domain.SimpleStatistics{
		Date:   "2002-03-30f",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	}

	tempMockStatistics := mockStatistics
	mockService := new(mocks.StatisticsService)

	j, err := json.Marshal(tempMockStatistics)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.SaveStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestSaveStatistics_ServiceError(t *testing.T) {
	mockStatistics := domain.SimpleStatistics{
		Date:   "2002-03-30",
		Views:  1,
		Clicks: 2,
		Cost:   3,
	}

	tempMockStatistics := mockStatistics
	mockService := new(mocks.StatisticsService)

	j, err := json.Marshal(tempMockStatistics)
	assert.NoError(t, err)

	mockService.On("SaveStatistics", mock.AnythingOfType("domain.Statistics")).Return(errors.New("Service Error"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.SaveStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}

func TestResetStatistics(t *testing.T) {
	mockService := new(mocks.StatisticsService)
	mockService.On("ResetStatistics").Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/ads", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.ResetStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockService.AssertExpectations(t)
}

func TestResetStatistics_ServiceError(t *testing.T) {
	mockService := new(mocks.StatisticsService)
	mockService.On("ResetStatistics").Return(errors.New("Service Error"))

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/ads", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := statisticsHttp.StatisticsHandler{
		StatisticsService: mockService,
	}
	err = handler.ResetStatistics(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}
