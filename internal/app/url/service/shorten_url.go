package service

import (
	"context"
	"time"
)

func (s *urlService) ShortenURL(
	ctx context.Context,
	originalURL string,
	expireIn int,
) (string, error) {

	// ✅ normalize expire time
	if expireIn <= 0 {
		expireIn = defaultExpireSeconds
	}

	ttl := time.Duration(expireIn) * time.Second

	// ✅ retry loop
	for i := 0; i < maxRetryAttempts; i++ {

		// generate short code
		code, err := s.gen.GenerateCode(defaultUrlCodeLength)
		if err != nil {
			return "", err
		}

		// try to store
		ok, err := s.repo.StoreURLIfNotExists(
			ctx,
			code,
			originalURL,
			ttl,
		)
		if err != nil {
			return "", err
		}

		// success
		if ok {
			return code, nil
		}
	}

	// exceeded retry
	return "", ErrMaxRetryExceeded
}
