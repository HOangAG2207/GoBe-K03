package redis

import (
	"fmt"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.RedisConfig, dbOverride *int) (*redis.Client, error) {
	db := cfg.DB
	if dbOverride != nil {
		db = *dbOverride
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       db,
	})

	return redisClient, nil
}
