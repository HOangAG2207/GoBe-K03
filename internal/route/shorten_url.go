package route

import (
	url_handler "github.com/HOangAG2207/GoBe-K03/internal/app/url/handler"
	url_repository "github.com/HOangAG2207/GoBe-K03/internal/app/url/repository"
	url_service "github.com/HOangAG2207/GoBe-K03/internal/app/url/service"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// RegisterPasswordRoutes groups all password-related endpoints
func RegisterUrlRoutes(api *echo.Group, redisClient *redis.Client) {

	// ===== Dependencies =====
	urlRepo := url_repository.NewUrlRepository(redisClient)
	urlService := url_service.NewUrlService(urlRepo, utils.NewCodeGenerator())
	urlHandler := url_handler.NewUrlHandler(urlService)

	// ===== Routes =====
	api.POST("/shorten-url", urlHandler.ShortURL)
}
