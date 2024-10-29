package app

import (
	"RunningTracker/internal/app/handlers"
	"RunningTracker/internal/app/repositories"
	"RunningTracker/internal/app/services"
	"RunningTracker/internal/infra/db"
)

type Dependencies struct {
	RunnerHandler *handlers.RunnerHandler
}

func SetupDependencies() *Dependencies {
	database := db.NewDataBaseConnection()

	runnerRepo := repositories.NewRunnerRepository(database)
	runnerService := services.NewRunnerService(runnerRepo)
	runnerHandler := handlers.NewRunnerHandler(runnerService)

	return &Dependencies{
		RunnerHandler: runnerHandler,
	}
}
