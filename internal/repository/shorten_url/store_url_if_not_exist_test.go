package shorten_url_repository

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HOangAG2207/GoBe-K03/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepository_StoreURLIfNotExists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		code     string
		url      string
		expireIn int

		setupMock func(ctx context.Context) *redis.Client

		expectedResult bool
		expectedError  error
	}{
		{
			name: "store new URL successfully",

			code:     "abc123",
			url:      "https://example.com",
			expireIn: 3600,

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				return redisClient
			},

			expectedResult: true,
			expectedError:  nil,
		},
		{
			name: "do not overwrite existing URL",

			code:     "abc123",
			url:      "https://example.com/updated",
			expireIn: 3600,

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.Set(ctx, "abc123", "https://example.com", time.Hour)
				return redisClient
			},

			expectedResult: false,
			expectedError:  nil,
		},
		{
			name: "failed to store URL due to closed Redis client",

			code:     "abc123",
			url:      "https://example.com",
			expireIn: 3600,

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.Close()
				return redisClient
			},

			expectedResult: false,
			expectedError:  redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			redisClient := tc.setupMock(ctx)
			repo := NewUrlRepository(redisClient)

			result, err := repo.StoreURLIfNotExists(ctx, tc.code, tc.url, tc.expireIn)
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
