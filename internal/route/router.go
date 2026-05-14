package route

import "github.com/labstack/echo/v4"

// RegisterRoutes is the central place to manage all routes
func RegisterRoutes(e *echo.Echo) {

	api := e.Group("/api")
	// ===== Register Password modules =====
	RegisterPasswordRoutes(api)
}
