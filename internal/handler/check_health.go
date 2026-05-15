package handler

import (
	"net/http"

	"github.com/HOangAG2207/GoBe-K03/internal/service"
	"github.com/labstack/echo/v4"
)

// HealthCheck defines the handler interface for health check endpoint
type HealthCheck interface {
	CheckHealth(c echo.Context) error
}

// healthCheckHandler implements HealthCheck interface
// It depends on service layer to get health information
type healthCheckHandler struct {
	service service.HealthCheck
}

// NewHealthCheck creates a new instance of HealthCheck handler
// using dependency injection for the service layer
func NewHealthCheck(s service.HealthCheck) HealthCheck {
	return &healthCheckHandler{
		service: s,
	}
}

// CheckHealth handles the HTTP request for health check endpoint
// It calls the service layer, gets the result, and returns JSON response
func (h *healthCheckHandler) CheckHealth(c echo.Context) error {
	// Call service layer to get health status
	res := h.service.CheckHealth()

	// Return HTTP 200 with JSON response
	return c.JSON(http.StatusOK, res)
}
