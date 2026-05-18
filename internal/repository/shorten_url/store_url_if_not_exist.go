package shorten_url_repository

import (
	"context"
	"time"
)

func (s *urlRepository) StoreURLIfNotExists(
	ctx context.Context,
	code, url string,
	exp time.Duration,
) (bool, error) {

	if exp <= 0 {
		exp = 60 * time.Second
	}

	ok, err := s.redisClient.SetNX(
		ctx,
		code,
		url,
		exp,
	).Result()

	if err != nil {
		return false, err
	}

	return ok, nil
}
