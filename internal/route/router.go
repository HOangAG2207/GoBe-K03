// internal/route/router.go
package route

import (
	"github.com/labstack/echo/v4"
)

type AppConfig struct {
	ServiceName string
	InstanceID  string
	// thêm biến khác ở đây...
}

// RegisterRoutes is the central place to manage all routes
func RegisterRoutes(e *echo.Echo, cfg AppConfig) {

	api := e.Group("/api")
	// ===== Register Health modules =====
	RegisterHealthRoutes(api, cfg)
	// ===== Register Password modules =====
	RegisterPasswordRoutes(api)
}
