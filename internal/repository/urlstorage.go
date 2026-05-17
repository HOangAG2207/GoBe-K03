package repository

import (
	"context"
	"errors"

	"github.com/HOangAG2207/GoBe-K03/pkg/redis"
)

const (
	urlExpiration = 60 * 60 // 1 hour
)

// Domain error
var ErrURLNotFound = errors.New("url not found")

// Interface
type URLStorage interface {
	StoreURL(ctx context.Context, code string, url string) error
	GetURL(ctx context.Context, code string) (string, error)
}

// Implementation
type urlStorage struct {
	redisClient *redis.Client
}

// Constructor
func NewURLStorage(redisClient *redis.Client) URLStorage {
	return &urlStorage{
		redisClient: redisClient,
	}
}

// Store short code -> original URL
func (s *urlStorage) StoreURL(ctx context.Context, code, url string) error {
	return s.redisClient.Set(ctx, code, url, urlExpiration)
}

// Get original URL
func (s *urlStorage) GetURL(ctx context.Context, code string) (string, error) {
	val, err := s.redisClient.Get(ctx, code)
	if err != nil {
		if errors.Is(err, redis.ErrKeyNotFound) {
			return "", ErrURLNotFound
		}
		return "", err
	}
	return val, nil
}
