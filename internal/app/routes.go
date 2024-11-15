package app

import (
	"RunningTracker/internal/app/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, runnerHandler *handlers.RunnerHandler) {
	// Grupo de rotas para o recurso Runner
	e.POST("/runners", runnerHandler.CreateRunner)
	e.PUT("/runners/:id", runnerHandler.UpdateRunner)
	e.DELETE("/runners/:id", runnerHandler.DeleteRunner)
}
