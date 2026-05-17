package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/redis/go-redis/v9"
)

// Domain errors
var (
	ErrKeyNotFound = errors.New("redis: key not found")
)

// Client wraps Redis client
type Client struct {
	rdb *redis.Client
}

// Constructor (không còn prefix)
func NewRedisClient(cfg config.RedisConfig, dbOverride *int) *Client {
	db := cfg.DB
	if dbOverride != nil {
		db = *dbOverride
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       db,
	})

	return &Client{
		rdb: rdb,
	}
}

//
// ===== PUBLIC METHODS =====
//

// Set key with TTL (seconds)
func (c *Client) Set(ctx context.Context, key string, value any, ttlSeconds int) error {
	return c.rdb.Set(
		ctx,
		key,
		value,
		time.Duration(ttlSeconds)*time.Second,
	).Err()
}

// Get value by key
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}
		return "", err
	}
	return val, nil
}

// Delete key
func (c *Client) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}
