package shorten_link_repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	urlExpiration = time.Hour
)

type UrlRepository interface {
	StoreUrl(ctx context.Context, shortCode string, originalUrl string) error
	GetUrl(ctx context.Context, shortCode string) (string, error)
}

type urlRepository struct {
	redisClient *redis.Client
}

func NewUrlRepository(redisClient *redis.Client) UrlRepository {
	return &urlRepository{redisClient: redisClient}
}
