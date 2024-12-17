package main

import (
	"RunningTracker/internal/app"
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/handlers"
	"RunningTracker/internal/app/repositories"
	"RunningTracker/internal/app/services"
	"RunningTracker/internal/infra/db"
	"RunningTracker/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//deps := app.SetupDependencies()
	//app.SetupRoutes(e, deps.RunnerHandler)

	db := db.NewDataBaseConnection()
	// Realiza a migração da entidade Runner
	db.AutoMigrate(&entities.Runner{})
	db.AutoMigrate(&entities.Race{})

	// Configuração das dependências
	runnerRepo := repositories.NewRunnerRepository(db)
	runnerService := services.NewRunnerService(runnerRepo)
	runnerHandler := handlers.NewRunnerHandler(runnerService)

	raceRepo := repositories.NewRaceRepository(db)
	raceService := services.NewRaceService(raceRepo)
	raceHandler := handlers.NewRaceHandler(raceService)

	// Rota para criar um novo corredor
	app.SetupDependencies()
	app.SetupRoutes(e, runnerHandler, raceHandler)

	e.Use(middleware.Logger)

	e.Start(":8080")
}
