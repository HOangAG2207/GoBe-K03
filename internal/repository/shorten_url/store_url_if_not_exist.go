package shorten_url_repository

import (
	"context"
	"time"
)

func (s *urlRepository) StoreURLIfNotExists(
	ctx context.Context,
	code, url string,
	exp int,
) (bool, error) {

	if exp <= 0 {
		exp = 60 // default fallback
	}

	ok, err := s.redisClient.SetNX(
		ctx,
		code,
		url,
		time.Duration(exp)*time.Second,
	).Result()

	if err != nil {
		return false, err
	}

	return ok, nil
}
