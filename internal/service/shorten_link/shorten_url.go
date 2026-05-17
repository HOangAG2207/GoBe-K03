package shorten_link

import (
	"context"
	"time"
)

func (s *urlService) ShortenURL(ctx context.Context, originalURL string, expireIn time.Duration) (string, error) {
	for range maxRetryAttempts {
		code, err := s.randomCodeGenerate.GenerateCode(defaultUrlCodeLength)
		if err != nil {
			return "", err
		}

		stored, err := s.repo.StoreURLIfNotExists(ctx, code, originalURL, expireIn)
		if err != nil {
			return "", err
		}
		if stored {
			return code, nil
		}
	}
	return "", ErrMaxRetryExceeded
}
