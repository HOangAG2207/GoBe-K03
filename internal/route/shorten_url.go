package route

import (
	shorten_url_handler "github.com/HOangAG2207/GoBe-K03/internal/handler/shorten_url"
	shorten_url_repository "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_url"
	shorten_url_service "github.com/HOangAG2207/GoBe-K03/internal/service/shorten_url"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// RegisterPasswordRoutes groups all password-related endpoints
func RegisterUrlRoutes(api *echo.Group, redisClient *redis.Client) {

	// ===== Dependencies =====
	urlRepo := shorten_url_repository.NewUrlRepository(redisClient)
	urlService := shorten_url_service.NewUrlService(urlRepo, utils.NewCodeGenerator())
	urlHandler := shorten_url_handler.NewUrlHandler(urlService)

	// ===== Routes =====
	api.POST("/shorten-url", urlHandler.ShortURL)
}
