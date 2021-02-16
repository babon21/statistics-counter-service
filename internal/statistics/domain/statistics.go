package domain

import "math"

type Statistics struct {
	SimpleStatistics
	Cpc float32 `json:"cpc"`
	Cpm float32 `json:"cpm"`
}

type SimpleStatistics struct {
	Date   string  `json:"date" validate:"required"`
	Views  int     `json:"views" validate:"gte=0"`
	Clicks int     `json:"clicks" validate:"gte=0"`
	Cost   float64 `json:"cost" validate:"gte=0"`
}

func calculateCpc(cost float64, clicks int) float32 {
	if clicks == 0 {
		return 0
	}
	cpc := cost / float64(clicks)
	return float32(math.Ceil(cpc*100) / 100)
}

func calculateCpm(cost float64, views int) float32 {
	if views == 0 {
		return 0
	}
	cpm := cost / float64(views)
	return float32(math.Ceil(cpm*100) / 100)
}

func NewStatistics(statistics SimpleStatistics) Statistics {
	statistics.Cost = math.Ceil(statistics.Cost*100) / 100
	return Statistics{
		SimpleStatistics: statistics,
		Cpc:              calculateCpc(statistics.Cost, statistics.Clicks),
		Cpm:              calculateCpm(statistics.Cost, statistics.Views),
	}
}
