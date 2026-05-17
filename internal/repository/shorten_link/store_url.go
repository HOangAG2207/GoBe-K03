package shorten_link

import (
	"context"
)

func (s *urlRepository) StoreUrl(ctx context.Context, shortCode string, originalUrl string) error {
	return s.redisClient.Set(ctx, shortCode, originalUrl, urlExpiration).Err()
}
