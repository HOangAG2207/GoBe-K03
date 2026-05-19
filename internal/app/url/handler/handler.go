package handler

import (
	"errors"

	url_service "github.com/HOangAG2207/GoBe-K03/internal/app/url/service"
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
	urlService url_service.Service
}

func NewUrlHandler(svc url_service.Service) Handler {
	return &urlHandler{
		urlService: svc,
	}
}
