package handlers

import (
	"RunningTracker/internal/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RaceHandler struct {
	service services.RaceService
}

func NewRaceHandler(service services.RaceService) *RaceHandler {
	return &RaceHandler{service: service}

}

func (h *RaceHandler) CreateRace(c echo.Context) error {
	var req struct {
		Distance      float64 `json:"distance"`
		Duration      float64 `json:"duration"`
		TypeOfRunning string  `json:"type_of_running"`
	}

	err := c.Bind(&req); if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	race, err := h.service.CreateRace(req.Distance, req.Duration, req.TypeOfRunning)
}
