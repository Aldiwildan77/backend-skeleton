package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// Endpoints are the entry point for route
func Endpoints(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running.")
	})
}
