package main

import (
	"fmt"
	"github.com/babon21/statistics-counter-service/internal/config"
	statisticsHttp "github.com/babon21/statistics-counter-service/internal/statistics/delivery/http"
	statisticsRepository "github.com/babon21/statistics-counter-service/internal/statistics/repository/postgres"
	statisticsService "github.com/babon21/statistics-counter-service/internal/statistics/service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.Init()

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", conf.Database.Username,
		conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	e := echo.New()
	//middL := middleware.InitMiddleware()
	//e.Use(middL.AccessLogMiddleware)
	roomRepo := statisticsRepository.NewPostgresStatisticsRepository(db)
	roomUsecase := statisticsService.NewStatisticsService(roomRepo)
	statisticsHttp.NewStatisticsHandler(e, roomUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
