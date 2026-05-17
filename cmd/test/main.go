package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	shorten_link_repository "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_link"
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

	shortCode := "gg:short"
	originalURL := "https://google.com"

	// 1. Store URL
	if err := urlRepo.StoreUrl(ctx, shortCode, originalURL); err != nil {
		log.Fatalf("store error: %v", err)
	}

	// 2. Get URL
	url, err := urlRepo.GetUrl(ctx, shortCode)
	if err != nil {
		log.Fatalf("get error: %v", err)
	}

	fmt.Println("Original URL:", url)
}
