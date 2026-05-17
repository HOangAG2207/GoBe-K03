package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/HOangAG2207/GoBe-K03/internal/repository"
	"github.com/HOangAG2207/GoBe-K03/pkg/redis"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	// ===== INIT REDIS =====
	rdb := redis.NewRedisClient(cfg.Redis, nil)

	// ===== INIT REPOSITORY =====
	urlRepo := repository.NewURLStorage(rdb)

	// ===== USE CASE =====

	// 1. Store URL
	err := urlRepo.StoreURL(ctx, "gg:short", "https://google.com")
	if err != nil {
		log.Fatalf("store error: %v", err)
	}

	// 2. Get URL
	url, err := urlRepo.GetURL(ctx, "gg:short")
	if err != nil {
		if err == repository.ErrURLNotFound {
			log.Println("URL not found")
			return
		}
		log.Fatalf("get error: %v", err)
	}

	fmt.Println("Original URL:", url)
}
