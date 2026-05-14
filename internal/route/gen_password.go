package route

import (
	"github.com/HOangAG2207/GoBe-K03/internal/handler"
	"github.com/HOangAG2207/GoBe-K03/internal/service"
	"github.com/labstack/echo/v4"
)

// RegisterPasswordRoutes groups all password-related endpoints
func RegisterPasswordRoutes(api *echo.Group) {

	// ===== Dependencies =====
	passwordService := service.NewGenPassword()
	passwordHandler := handler.NewGenPassword(passwordService)

	// ===== Routes =====
	api.GET("/gen-password", passwordHandler.GeneratePassword)
}
