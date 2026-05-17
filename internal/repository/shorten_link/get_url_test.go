package shorten_link

import (
	"context"
	"testing"

	redisPkg "github.com/HOangAG2207/GoBe-K03/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetUrl(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func() *redis.Client

		expectedURL   string
		expectedError error
	}{
		{
			name: "success - get url",
			setupMock: func() *redis.Client {
				client := redisPkg.InitMockRedis(t)

				err := client.Set(context.Background(), "test", "https://example.com", 0).Err()
				assert.Nil(t, err)

				return client
			},
			expectedURL:   "https://example.com",
			expectedError: nil,
		},
		{
			name: "fail - url not found",
			setupMock: func() *redis.Client {
				return redisPkg.InitMockRedis(t)
			},
			expectedURL:   "",
			expectedError: redis.Nil, // 👈 FIX
		},
		{
			name: "fail - redis connection closed",
			setupMock: func() *redis.Client {
				client := redisPkg.InitMockRedis(t)
				client.Close()
				return client
			},
			expectedURL:   "",
			expectedError: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := t.Context()

			redisClient := tc.setupMock()

			urlStorage := NewUrlRepository(redisClient)

			result, err := urlStorage.GetUrl(ctx, "test")

			assert.Equal(t, tc.expectedURL, result)

			if tc.expectedError != nil {
				assert.Error(t, err)

				if tc.expectedError == redis.Nil {
					assert.ErrorIs(t, err, redis.Nil)
				} else {
					assert.EqualError(t, err, tc.expectedError.Error())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
