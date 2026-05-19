// internal/route/router.go
package route

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type AppConfig struct {
	ServiceName string
	InstanceID  string
	// thêm biến khác ở đây...
}

// RegisterRoutes is the central place to manage all routes
func RegisterRoutes(e *echo.Echo, cfg AppConfig, redisClient *redis.Client) {

	api := e.Group("/api")

	// ===== Register Password modules =====
	RegisterPasswordRoutes(api)

	// ===== Register Health modules =====
	RegisterHealthRoutes(api, cfg)

	// ===== Register Url modules =====
	RegisterUrlRoutes(api, redisClient)
}
