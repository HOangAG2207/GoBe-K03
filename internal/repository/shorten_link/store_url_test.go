package shorten_link_repository

import (
	"context"
	"testing"

	redisPkg "github.com/HOangAG2207/GoBe-K03/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepository_StoreUrl(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func() *redis.Client

		expectedError error

		verifyFunc func(ctx context.Context, redisClient *redis.Client)
	}{
		{
			name: "success - store url",
			setupMock: func() *redis.Client {
				client := redisPkg.InitMockRedis(t)
				return client
			},
			expectedError: nil,
			verifyFunc: func(ctx context.Context, redisClient *redis.Client) {
				url, err := redisClient.Get(ctx, "test").Result()
				assert.Nil(t, err)
				assert.Equal(t, "https://example.com", url)
			},
		},
		{
			name: "fail - redis connection error",
			setupMock: func() *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.Close()
				return redisClient
			},

			expectedError: redis.ErrClosed,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisMockClient := tc.setupMock()

			urlStorage := NewUrlRepository(redisMockClient)

			err := urlStorage.StoreUrl(ctx, "test", "https://example.com")
			assert.Equal(t, tc.expectedError, err)

			if err == nil {
				tc.verifyFunc(ctx, redisMockClient)
			}

		})
	}
}
