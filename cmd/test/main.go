package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	shorten_url_repository "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_url"

	shorten_url_service "github.com/HOangAG2207/GoBe-K03/internal/service/shorten_url"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
	redisPkg "github.com/HOangAG2207/GoBe-K03/pkg/redis"
)

func main() {
	ctx := context.Background()

	// ===== CONFIG =====
	cfg := config.Load()

	// ===== REDIS =====
	rdb, err := redisPkg.NewRedisClient(redisPkg.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}, nil)
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis not connected: %v", err)
	}

	defer rdb.Close()

	// ===== REPOSITORY =====
	urlRepo := shorten_url_repository.NewUrlRepository(rdb)

	// ===== SERVICE =====
	urlService := shorten_url_service.NewUrlService(
		urlRepo,
		utils.NewCodeGenerator(),
	)

	// ===== TEST RUN =====
	key, _ := urlService.ShortenURL(ctx, "https://google.com", 1)
	fmt.Println(key)
}
