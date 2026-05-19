package service

import (
	"context"
	"errors"
	"testing"
	"time"

	mockRepo "github.com/HOangAG2207/GoBe-K03/internal/app/url/repository/mocks"
	mockGen "github.com/HOangAG2207/GoBe-K03/internal/utils/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	errGen  = errors.New("gen error")
	errRepo = errors.New("repo error")
)

func TestService_ShortenURL(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := []struct {
		name string

		setupGen  func() *mockGen.CodeGenerator
		setupRepo func() *mockRepo.Repository

		expireIn int

		expectedCode string
		expectedErr  error
	}{
		{
			name: "success first try",

			setupGen: func() *mockGen.CodeGenerator {
				m := &mockGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("abc12345", nil).Once()
				return m
			},

			setupRepo: func() *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists",
					ctx,
					"abc12345",
					"https://test.com",
					time.Second,
				).Return(true, nil).Once()
				return m
			},

			expireIn:     1,
			expectedCode: "abc12345",
		},
		{
			name: "retry then success",

			setupGen: func() *mockGen.CodeGenerator {
				m := &mockGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).Return("dup", nil).Once()
				m.On("GenerateCode", defaultUrlCodeLength).Return("new", nil).Once()
				return m
			},

			setupRepo: func() *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists", ctx, "dup", "https://test.com", time.Second).
					Return(false, nil).Once()
				m.On("StoreURLIfNotExists", ctx, "new", "https://test.com", time.Second).
					Return(true, nil).Once()
				return m
			},

			expireIn:     1,
			expectedCode: "new",
		},
		{
			name: "generator error",

			setupGen: func() *mockGen.CodeGenerator {
				m := &mockGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("", errGen).Once()
				return m
			},

			setupRepo: func() *mockRepo.Repository {
				return &mockRepo.Repository{}
			},

			expireIn:    1,
			expectedErr: errGen,
		},
		{
			name: "repo error",

			setupGen: func() *mockGen.CodeGenerator {
				m := &mockGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("abc", nil).Once()
				return m
			},

			setupRepo: func() *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists", ctx, "abc", "https://test.com", time.Second).
					Return(false, errRepo).Once()
				return m
			},

			expireIn:    1,
			expectedErr: errRepo,
		},
		{
			name: "max retry exceeded",

			setupGen: func() *mockGen.CodeGenerator {
				m := &mockGen.CodeGenerator{}
				for i := 0; i < maxRetryAttempts; i++ {
					m.On("GenerateCode", defaultUrlCodeLength).
						Return("dup", nil).Once()
				}
				return m
			},

			setupRepo: func() *mockRepo.Repository {
				m := &mockRepo.Repository{}
				for i := 0; i < maxRetryAttempts; i++ {
					m.On("StoreURLIfNotExists", ctx, "dup", "https://test.com", time.Second).
						Return(false, nil).Once()
				}
				return m
			},

			expireIn:    1,
			expectedErr: ErrMaxRetryExceeded,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gen := tc.setupGen()
			repo := tc.setupRepo()

			s := NewUrlService(repo, gen)

			code, err := s.ShortenURL(ctx, "https://test.com", tc.expireIn)

			assert.Equal(t, tc.expectedCode, code)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error()) // ✅ FIX CHÍNH
			} else {
				assert.NoError(t, err)
			}

			gen.AssertExpectations(t)
			repo.AssertExpectations(t)
		})
	}
}
