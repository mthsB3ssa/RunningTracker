package handlers

import (
	"RunningTracker/internal/app/entities"
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

func (h *RunnerHandler) GetUsers(c echo.Context) error {
	runner, err := h.service.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, runner)
}

func (h *RunnerHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	runner, err := h.service.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, runner)
}

func (h *RunnerHandler) UpdateRunner(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Lê os dados do JSON do body
	var req struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	runner := &entities.Runner{
		ID:   id,
		Name: req.Name,
		Age:  req.Age,
	}

	updateRunner, err := h.service.UpdateRunner(runner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updateRunner)
}

func (h *RunnerHandler) DeleteRunner(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteRunner(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RunnerHandler) GetAllRunsByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	runners, err := h.service.GetAllRunsByUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, runners)
}
