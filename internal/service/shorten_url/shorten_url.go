package shorten_url_service

import (
	"context"
	"time"
)

func (s *urlService) ShortenURL(
	ctx context.Context,
	originalURL string,
	expireIn int,
) (string, error) {

	if expireIn <= 0 {
		expireIn = 60
	}

	ttl := time.Duration(expireIn) * time.Second

	for i := 0; i < maxRetryAttempts; i++ {

		code, err := s.randomCodeGenerate.GenerateCode(defaultUrlCodeLength)
		if err != nil {
			return "", err
		}

		stored, err := s.repo.StoreURLIfNotExists(
			ctx,
			code,
			originalURL,
			ttl, // ✅ FIX HERE
		)
		if err != nil {
			return "", err
		}

		if stored {
			return code, nil
		}
	}

	return "", ErrMaxRetryExceeded
}
