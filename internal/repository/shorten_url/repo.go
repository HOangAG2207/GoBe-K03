package shorten_url_repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	urlExpiration = time.Hour
)

//go:generate mockery --name Repository --filename shorten_url_mock.go --output ./mocks
type Repository interface {
	StoreUrl(ctx context.Context, shortCode string, originalUrl string) error
	GetUrl(ctx context.Context, shortCode string) (string, error)
	StoreURLIfNotExists(ctx context.Context, code, url string, exp int) (bool, error)
}

type urlRepository struct {
	redisClient *redis.Client
}

func NewUrlRepository(redisClient *redis.Client) Repository {
	return &urlRepository{redisClient: redisClient}
}
