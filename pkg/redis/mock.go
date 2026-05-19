package pkg_redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// Mock Redis chạy in-memory
func InitMockRedis(t *testing.T) *redis.Client {
	t.Helper()

	mock := miniredis.RunT(t)

	return redis.NewClient(&redis.Options{
		Addr: mock.Addr(),
	})
}
