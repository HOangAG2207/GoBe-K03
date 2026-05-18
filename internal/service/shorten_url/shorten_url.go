package shorten_url_service

import (
	"context"
)

func (s *urlService) ShortenURL(ctx context.Context, originalURL string, expireIn int) (string, error) {
	for i := 0; i < maxRetryAttempts; i++ {

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
