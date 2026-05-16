package main

import (
	"context"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/HOangAG2207/GoBe-K03/pkg/redis"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	// ✅ dùng default (ENV)
	rdb0 := redis.NewRedisClient(cfg.Redis, nil, nil)
	// set
	_ = rdb0.Set(ctx, "user:1", "hello", 60)

	// ------------------------

	// ✅ override DB
	db := 1
	rdb1 := redis.NewRedisClient(cfg.Redis, &db, nil)
	// set
	_ = rdb1.Set(ctx, "user:1", "hello", 60)

	// ✅ override prefix
	prefix := "user"
	rdb2 := redis.NewRedisClient(cfg.Redis, nil, &prefix)
	// set
	_ = rdb2.Set(ctx, "1", "hello", 60)
}
