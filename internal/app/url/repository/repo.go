package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	urlExpiration = time.Hour
)

//go:generate mockery --name Repository --filename url_repository_mock.go --output ./mocks
type Repository interface {
	StoreUrl(ctx context.Context, shortCode string, originalUrl string) error
	GetUrl(ctx context.Context, shortCode string) (string, error)
	StoreURLIfNotExists(ctx context.Context, code, url string, exp time.Duration) (bool, error)
}

type urlRepository struct {
	redisClient *redis.Client
}

func NewUrlRepository(redisClient *redis.Client) Repository {
	return &urlRepository{redisClient: redisClient}
}
func (s *urlRepository) StoreUrl(ctx context.Context, shortCode string, originalUrl string) error {
	return s.redisClient.Set(ctx, shortCode, originalUrl, urlExpiration).Err()
}

func (s *urlRepository) GetUrl(ctx context.Context, shortCode string) (string, error) {
	return s.redisClient.Get(ctx, shortCode).Result()
}
func (s *urlRepository) StoreURLIfNotExists(
	ctx context.Context,
	code, url string,
	exp time.Duration,
) (bool, error) {
	return s.redisClient.SetNX(
		ctx,
		code,
		url,
		exp,
	).Result()
}
