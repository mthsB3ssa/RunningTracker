package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

// Logger its a middleware that logs the request method, URL and response status
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		log.Printf("\n %s %s %d", req.Method, req.URL, res.Status)
		return next(c)
	}
}
