package handlers

import (
	"RunningTracker/internal/app/services"

	"github.com/labstack/echo/v4"
)

type RaceHandler struct {
	service services.RaceService
}

func NewRaceHandler(service services.RaceService) *RaceHandler {
	return &RaceHandler{service: service}

}

func (h *RaceHandler) CreateRace(c echo.Context) error {

}
