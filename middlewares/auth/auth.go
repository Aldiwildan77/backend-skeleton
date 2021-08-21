package auth

import (
	"github.com/Aldiwildan77/backend-skeleton/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JWT returns JWT auth middleware
func JWT() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(config.Cfg.JWTSecret),
	}

	return middleware.JWTWithConfig(config)
}
