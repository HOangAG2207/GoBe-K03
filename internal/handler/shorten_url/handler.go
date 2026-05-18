package shorten_url_handler

import (
	"errors"

	shorten_url_service "github.com/HOangAG2207/GoBe-K03/internal/service/shorten_url"
	"github.com/labstack/echo/v4"
)

var (
	ErrCodeNotFound       = errors.New("code not found")
	InValidRequestPayload = errors.New("invalid request payload")
	InternalServerError   = errors.New("internal server error")
)

type Handler interface {
	ShortURL(c echo.Context) error
}

type urlHandler struct {
	urlService shorten_url_service.Service
}

func NewUrlHandler(svc shorten_url_service.Service) Handler {
	return &urlHandler{
		urlService: svc,
	}
}
