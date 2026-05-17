package shorten_link

import (
	"context"
	"time"
)

func (s *urlRepository) StoreURLIfNotExists(ctx context.Context, code, url string, exp time.Duration) (bool, error) {
	return s.redisClient.SetNX(ctx, code, url, exp).Result()
}
