package handlers

import (
	"RunningTracker/internal/app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Armezena uma instância do RunnerService
type RunnerHandler struct {
	service services.RunnerService
}

func NewRunnerHandler(service services.RunnerService) *RunnerHandler {
	return &RunnerHandler{service: service}
}

/*
	Método que recebe a requisição HTTP

*/
func (h *RunnerHandler) CreateRunner(c echo.Context) error {
	var req struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	runner, err := h.service.CreateRunner(req.Name, req.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, runner)
}

func (h *RunnerHandler) UpdateRunner(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.ParseUint(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updateRunner, err := h.service.UpdateRunner(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updateRunner)
}

func (h *RunnerHandler) DeleteRunner(c echo.Context) error {
}
^