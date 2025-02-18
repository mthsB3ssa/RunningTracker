package app

import (
	"RunningTracker/internal/app/controllers"
	"RunningTracker/internal/app/repositories"
	"RunningTracker/internal/app/services"
	"RunningTracker/internal/infra/db"

	"gorm.io/gorm"
)

type Dependencies struct {
	DB            *gorm.DB
	RunnerHandler *handlers.RunnerHandler
	RaceHandler   *handlers.RaceHandler
}

func SetupDependencies() (*Dependencies, error) {
	database := db.NewDataBaseConnection()

	runnerRepo := repositories.NewRunnerRepository(database)
	runnerService := services.NewRunnerService(runnerRepo)
	runnerHandler := handlers.NewRunnerHandler(runnerService)

	raceRepo := repositories.NewRaceRepository(database)
	raceService := services.NewRaceService(raceRepo)
	raceHandler := handlers.NewRaceHandler(raceService)

	return &Dependencies{
		RunnerHandler: runnerHandler,
		RaceHandler:   raceHandler,
		DB:            database,
	}, nil
}
