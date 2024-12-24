package main

import (
	"RunningTracker/internal/app"
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/infra/db"
	"RunningTracker/pkg/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := db.NewDataBaseConnection()
	// Realiza a migração da entidade Runner
	db.AutoMigrate(&entities.Runner{})
	db.AutoMigrate(&entities.Race{})

	// Rota para criar um novo corredor
	dependencies, err := app.SetupDependencies()
	if err != nil {
		log.Fatal("Error in intitialize the dependencies", err)
	}

	app.SetupRoutes(e, dependencies.RunnerHandler, dependencies.RaceHandler)

	e.Use(middleware.Logger)

	e.Start(":8080")
}
