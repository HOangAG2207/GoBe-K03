package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg Config, dbOverride *int) (*redis.Client, error) {
	db := cfg.DB
	if dbOverride != nil {
		db = *dbOverride
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       db,
	})

	return client, nil
}
