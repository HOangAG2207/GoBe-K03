// package pkg_redis

// import (
// 	"fmt"

// 	"github.com/redis/go-redis/v9"
// )

// func NewRedisClient(cfg Config, dbOverride *int) (*redis.Client, error) {
// 	db := cfg.DB
// 	if dbOverride != nil {
// 		db = *dbOverride
// 	}

// 	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

// 	client := redis.NewClient(&redis.Options{
// 		Addr:     addr,
// 		Password: cfg.Password,
// 		DB:       db,
// 	})

//		return client, nil
//	}
package pkg_redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Client struct {
	rdb *redis.Client
}

// Constructor
func NewClient(cfg Config) *Client {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Client{rdb: rdb}
}

// Optional: override DB (nếu cần multi-tenant)
func NewClientWithDB(cfg Config, db int) *Client {
	cfg.DB = db
	return NewClient(cfg)
}

// Health check
func (c *Client) Ping(ctx context.Context) error {
	_, err := c.rdb.Ping(ctx).Result()
	return err
}
