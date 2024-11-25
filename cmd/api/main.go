package main

import (
	"RunningTracker/internal/app"
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/handlers"
	"RunningTracker/internal/app/repositories"
	"RunningTracker/internal/app/services"
	"RunningTracker/pkg/middleware"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	//deps := app.SetupDependencies()
	//app.SetupRoutes(e, deps.RunnerHandler)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Realiza a migração da entidade Runner
	db.AutoMigrate(&entities.Runner{})

	// Configuração das dependências
	runnerRepo := repositories.NewRunnerRepository(db)
	runnerService := services.NewRunnerService(runnerRepo)
	runnerHandler := handlers.NewRunnerHandler(runnerService)

	// Rota para criar um novo corredor
	app.SetupRoutes(e, runnerHandler)

	e.Use(middleware.Logger)

	e.Start(":8080")
}
