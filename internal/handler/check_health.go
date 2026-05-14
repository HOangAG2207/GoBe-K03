package handler

import (
	"net/http"

	"github.com/HOangAG2207/GoBe-K03/internal/service"
	"github.com/labstack/echo/v4"
)

type HealthCheck interface {
	CheckHealth(c echo.Context) error
}
type healthCheckHandler struct {
	service service.HealthCheck
}

func NewHealthCheck(s service.HealthCheck) HealthCheck {
	return &healthCheckHandler{
		service: s,
	}
}
func (h *healthCheckHandler) CheckHealth(c echo.Context) error {
	res := h.service.CheckHealth()
	return c.JSON(http.StatusOK, res)
}
