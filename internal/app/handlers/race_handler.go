package handlers

import (
	"RunningTracker/internal/app/services"
	"net/http"
	"strconv"

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
		RunnerID      int     `json:"runner_id"`
		Distance      float64 `json:"distance"`
		Duration      float64 `json:"duration"`
		TypeOfRunning string  `json:"type_of_running"`
	}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	race, err := h.service.CreateRace(req.RunnerID, req.Distance, req.Duration, req.TypeOfRunning)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, race)
}

func (h *RaceHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	race, err := h.service.FindById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusFound, race)
}
