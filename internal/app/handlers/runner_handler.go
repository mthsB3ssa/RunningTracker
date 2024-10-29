package handlers

import (
	"RunningTracker/internal/app/services"
	"net/http"

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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid output"})
	}

	runner, err := h.service.CreateRunner(req.Name, req.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, runner)
}
