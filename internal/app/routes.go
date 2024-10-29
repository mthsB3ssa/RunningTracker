package app

import (
	"RunningTracker/internal/app/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, runnerHandler *handlers.RunnerHandler) {
	// Grupo de rotsas para o recurso Runner
	e.POST("/runner", runnerHandler.CreateRunner)
}