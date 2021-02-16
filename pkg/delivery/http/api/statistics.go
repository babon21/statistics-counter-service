package api

import "github.com/babon21/statistics-counter-service/internal/statistics/domain"

type SaveStatisticsRequest struct {
	domain.SimpleStatistics
}

type GetStatisticsRequest struct {
	SortField string `json:"sort_by"`
	SortOrder string `json:"order_by"`
	From      string `json:"from"`
	To        string `json:"to"`
}

type GetStatisticsResponse struct {
	Statistics []domain.Statistics `json:"statistics"`
}
