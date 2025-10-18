package router

import (
	"github.com/krabhinav07/homefixit/internal/handler"

	"github.com/labstack/echo/v4"
)

func registerSystemRoutes(r *echo.Echo, h *handler.Handlers) {
	r.GET("/status", h.Health.CheckHealth)

	r.Static("/static", "static")

	r.GET("/docs", h.OpenAPI.ServeOpenAPIUI)
}
