package app

import (
	"RunningTracker/internal/app/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, runnerHandler *handlers.RunnerHandler) {
	// Grupo de rotas para o recurso Runner
	e.POST("/runners", runnerHandler.CreateRunner)
	e.GET("/runners/:id", runnerHandler.GetRunner)
	e.DELETE("/runners/:id", runnerHandler.DeleteRunner)
}