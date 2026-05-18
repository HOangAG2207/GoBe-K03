package shorten_url_service

import (
	"context"
	"errors"

	shorten_url "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_url"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
)

const (
	defaultUrlCodeLength = 8
	maxRetryAttempts     = 10
)

var (
	ErrCodeNotFound     = errors.New("short code not found")
	ErrMaxRetryExceeded = errors.New("maximum retry attempts exceeded")
)

//go:generate mockery --name Service --filename shorten_url_mock.go --output ./mocks
type Service interface {
	ShortenURL(ctx context.Context, originalURL string, expireIn int) (string, error)
}

type urlService struct {
	repo               shorten_url.Repository
	randomCodeGenerate utils.CodeGenerator
}

func NewUrlService(repo shorten_url.Repository, randomCodeGenerate utils.CodeGenerator) Service {
	return &urlService{repo: repo, randomCodeGenerate: randomCodeGenerate}
}
