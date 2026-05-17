package shorten_link

import "context"

func (s *urlRepository) GetUrl(ctx context.Context, shortCode string) (string, error) {
	return s.redisClient.Get(ctx, shortCode).Result()
}
