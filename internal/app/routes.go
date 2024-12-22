package app

import (
	"RunningTracker/internal/app/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, runnerHandler *handlers.RunnerHandler, raceHandler *handlers.RaceHandler) {
	// Grupo de rotas para o recurso Runner
	e.POST("/runners", runnerHandler.CreateRunner)
	e.GET("/runnners", runnerHandler.GetUsers)
	e.GET("/runners/:id", runnerHandler.FindById)
	e.PUT("/runners/:id", runnerHandler.UpdateRunner)
	e.DELETE("/runners/:id", runnerHandler.DeleteRunner)

	//Group of routes to the Race
	e.POST("/races", raceHandler.CreateRace)
	e.GET("/races/:id", raceHandler.FindById)
	e.DELETE("/races/:id", raceHandler.FindById)
}
