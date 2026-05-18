package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	shorten_link_repository "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_link"
	shorten_link_service "github.com/HOangAG2207/GoBe-K03/internal/service/shorten_link"
	"github.com/HOangAG2207/GoBe-K03/internal/utils"
	redisPkg "github.com/HOangAG2207/GoBe-K03/pkg/redis"
)

func main() {
	ctx := context.Background()

	// ===== LOAD CONFIG =====
	cfg := config.Load()

	// ===== INIT REDIS =====
	rdb, err := redisPkg.NewRedisClient(cfg.Redis, nil)
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	// (optional nhưng nên có)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis not connected: %v", err)
	}

	defer func() {
		if err := rdb.Close(); err != nil {
			log.Println("redis close error:", err)
		}
	}()

	// ===== INIT REPOSITORY =====
	urlRepo := shorten_link_repository.NewUrlRepository(rdb)

	// ===== USE CASE =====

	originalURL := "https://google.com"

	urlService := shorten_link_service.NewUrlService(urlRepo, utils.NewCodeGenerator())

	key, _ := urlService.ShortenURL(ctx, originalURL, time.Hour)
	fmt.Println(key)
}
