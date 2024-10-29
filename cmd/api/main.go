package main

import (
	"RunningTracker/internal/app"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	deps := app.SetupDependencies()

	app.SetupRoutes(e, deps.RunnerHandler)

	e.Start(":8080")
}
