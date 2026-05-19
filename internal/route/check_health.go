// internal/route/check_health.go
package route

import (
	"github.com/HOangAG2207/GoBe-K03/internal/app/health/handler"
	"github.com/HOangAG2207/GoBe-K03/internal/app/health/service"
	"github.com/labstack/echo/v4"
)

// RegisterHealthRoutes groups all health-related endpoints
func RegisterHealthRoutes(egr *echo.Group, cfg AppConfig) {

	// ===== Dependencies =====
	healthService := service.NewHealthCheck(cfg.ServiceName, cfg.InstanceID)
	healthHandler := handler.NewHealthCheck(healthService)

	// ===== Routes =====
	egr.GET("/health-check", healthHandler.CheckHealth)
}
