package main

import (
	"RunningTracker/internal/app"
	"RunningTracker/internal/app/entities"
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

	// Rota para criar um novo corredor
	dependencies := app.SetupDependencies()
	app.SetupRoutes(e, dependencies.RunnerHandler, dependencies.RaceHandler)

	e.Use(middleware.Logger)

	e.Start(":8080")
}
