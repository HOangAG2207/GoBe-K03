package service

import (
	"context"
	"errors"

	shorten_url "github.com/HOangAG2207/GoBe-K03/internal/app/url/repository"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
)

const (
	defaultUrlCodeLength = 8
	maxRetryAttempts     = 10
	defaultExpireSeconds = 60
)

var (
	ErrCodeNotFound     = errors.New("short code not found")
	ErrMaxRetryExceeded = errors.New("maximum retry attempts exceeded")
)

//go:generate mockery --name Service --filename url_service_mock.go --output ./mocks
type Service interface {
	ShortenURL(ctx context.Context, originalURL string, expireIn int) (string, error)
}

type urlService struct {
	repo shorten_url.Repository
	gen  utils.CodeGenerator
}

func NewUrlService(
	repo shorten_url.Repository,
	gen utils.CodeGenerator,
) Service {
	return &urlService{
		repo: repo,
		gen:  gen,
	}
}
