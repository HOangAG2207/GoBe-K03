package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/redis/go-redis/v9"
)

// Client wraps the Redis client and adds key prefix support
type Client struct {
	rdb    *redis.Client // underlying Redis client
	prefix string        // key prefix for namespacing
}

// NewRedisClient initializes a new Redis client instance
// It supports overriding DB index and key prefix dynamically
func NewRedisClient(cfg config.RedisConfig, dbOverride *int, prefixOverride *string) *Client {
	// Use default DB from config
	db := cfg.DB

	// Override DB if provided
	if dbOverride != nil {
		db = *dbOverride
	}

	// Use default prefix from config
	prefix := cfg.Prefix

	// Override prefix if provided
	if prefixOverride != nil {
		prefix = *prefixOverride
	}

	// Build Redis address (host:port)
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	// Create Redis client with given configuration
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       db,
	})

	return &Client{
		rdb:    rdb,
		prefix: prefix,
	}
}

// buildKey applies prefix to the given key if prefix is set
// Example: prefix="app" and key="user:1" -> "app:user:1"
func (c *Client) buildKey(key string) string {
	if c.prefix == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", c.prefix, key)
}

// Set stores a key-value pair in Redis with a TTL (in seconds)
// ttlSeconds = 0 means no expiration
func (c *Client) Set(ctx context.Context, key string, value any, ttlSeconds int) error {
	return c.rdb.Set(
		ctx,
		c.buildKey(key),
		value,
		time.Duration(ttlSeconds)*time.Second,
	).Err()
}
