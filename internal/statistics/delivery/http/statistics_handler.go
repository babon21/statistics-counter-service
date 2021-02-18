package http

import (
	"errors"
	"fmt"
	"github.com/babon21/statistics-counter-service/internal/statistics/domain"
	"github.com/babon21/statistics-counter-service/internal/statistics/service"
	"github.com/babon21/statistics-counter-service/pkg/delivery/http/api"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// StatisticsHandler  represent the httphandler for statistics
type StatisticsHandler struct {
	StatisticsService service.StatisticsService
}

func (h *StatisticsHandler) GetStatisticsList(c echo.Context) error {
	var request api.GetStatisticsRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	err = validateRequestDates(request.From, request.To)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	sortField, sortOrder, err := parseSortParams(request.SortField, request.SortOrder)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	statistics, err := h.StatisticsService.GetStatisticsList(request.From, request.To, sortField, sortOrder)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, ResponseError{Message: err.Error()}, "  ")
	}

	response := api.GetStatisticsResponse{Statistics: statistics}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func validateRequestDates(from string, to string) error {
	err := validateDate(from)
	if err != nil {
		return errors.New("Invalid date format 'from' field")
	}

	err = validateDate(to)
	if err != nil {
		return errors.New("Invalid date format 'to' field")
	}

	return nil
}

func parseSortParams(sortFieldParam string, sortOrderParam string) (service.SortField, service.SortOrder, error) {
	if sortFieldParam == "" {
		return "", "", SortByParamIsEmpty
	}

	if sortOrderParam == "" {
		return "", "", OrderByParamIsEmpty
	}

	sortOrder := service.SortOrder(sortOrderParam)
	if sortOrder != service.AscOrder && sortOrder != service.DescOrder {
		return "", "", WrongOrderByParam
	}

	sortField := service.SortField(sortFieldParam)
	switch sortField {
	case service.DateField, service.ViewsField, service.ClicksField, service.CostField, service.CpcField, service.CpmField:
		return sortField, sortOrder, nil
	default:
		return "", "", WrongSortField
	}
}

func (h *StatisticsHandler) SaveStatistics(c echo.Context) error {
	var request api.SaveStatisticsRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	err = validateSaveStatisticsRequest(request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	statistics := domain.NewStatistics(domain.SimpleStatistics{
		Date:   request.Date,
		Views:  request.Views,
		Clicks: request.Clicks,
		Cost:   request.Cost,
	})

	err = h.StatisticsService.SaveStatistics(statistics)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, ResponseError{Message: err.Error()}, "  ")
	}

	return c.NoContent(http.StatusCreated)
}

func validateSaveStatisticsRequest(request api.SaveStatisticsRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		var errMessage string
		for _, err := range err.(validator.ValidationErrors) {
			errMessage = fmt.Sprintf("%s must %s %s, but actual value %v", err.Field(), err.Tag(), err.Param(), err.Value())
		}
		return errors.New(errMessage)
	}

	return validateDate(request.Date)
}

func validateDate(date string) error {
	_, err := time.Parse(service.LayoutISO, date)
	return err
}

func (h *StatisticsHandler) ResetStatistics(c echo.Context) error {
	if err := h.StatisticsService.ResetStatistics(); err != nil {
		return c.JSONPretty(http.StatusInternalServerError, ResponseError{Message: err.Error()}, "  ")
	}
	return c.NoContent(http.StatusNoContent)
}

// NewStatisticsHandler will initialize the statistics/ resources endpoint
func NewStatisticsHandler(e *echo.Echo, s service.StatisticsService) {
	handler := &StatisticsHandler{
		StatisticsService: s,
	}

	e.GET("/statistics", handler.GetStatisticsList)
	e.POST("/statistics", handler.SaveStatistics)
	e.DELETE("/statistics", handler.ResetStatistics)
}
