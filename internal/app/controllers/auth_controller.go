package handlers

import "github.com/labstack/echo/v4"

// Armezena uma instância do RunnerService
type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (h *AuthController) Login(c echo.Context) {
	
}
