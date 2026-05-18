package shorten_url_service

import (
	"context"
	"errors"
	"testing"
	"time"

	mockRepo "github.com/HOangAG2207/GoBe-K03/internal/repository/shorten_url/mocks"
	mockRandomCodeGen "github.com/HOangAG2207/GoBe-K03/internal/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestService_ShortenUrl(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockRandomCodeGen func() *mockRandomCodeGen.CodeGenerator
		setupMockRepo          func(ctx context.Context) *mockRepo.Repository

		inputOriginalURL string
		inputExpireIn    int

		expectedCode  string
		expectedError error
	}{
		{
			name: "success on first attempt",

			setupMockRandomCodeGen: func() *mockRandomCodeGen.CodeGenerator {
				m := &mockRandomCodeGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("abc12345", nil).Once()
				return m
			},
			setupMockRepo: func(ctx context.Context) *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists", ctx, "abc12345", "https://test.com", time.Minute).
					Return(true, nil).Once()
				return m
			},

			inputOriginalURL: "https://test.com",
			inputExpireIn:    1,

			expectedCode:  "abc12345",
			expectedError: nil,
		},
		{
			name: "retry then success",

			setupMockRandomCodeGen: func() *mockRandomCodeGen.CodeGenerator {
				m := &mockRandomCodeGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("dupcode", nil).Once()
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("newcode", nil).Once()
				return m
			},
			setupMockRepo: func(ctx context.Context) *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists", ctx, "dupcode", "https://test.com", time.Minute).
					Return(false, nil).Once()
				m.On("StoreURLIfNotExists", ctx, "newcode", "https://test.com", time.Minute).
					Return(true, nil).Once()
				return m
			},

			inputOriginalURL: "https://test.com",
			inputExpireIn:    1,

			expectedCode:  "newcode",
			expectedError: nil,
		},
		{
			name: "random generator error",

			setupMockRandomCodeGen: func() *mockRandomCodeGen.CodeGenerator {
				m := &mockRandomCodeGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("", errors.New("gen error")).Once()
				return m
			},
			setupMockRepo: func(ctx context.Context) *mockRepo.Repository {
				return &mockRepo.Repository{}
			},

			inputOriginalURL: "https://test.com",
			inputExpireIn:    1,

			expectedCode:  "",
			expectedError: errors.New("gen error"),
		},
		{
			name: "repository error",

			setupMockRandomCodeGen: func() *mockRandomCodeGen.CodeGenerator {
				m := &mockRandomCodeGen.CodeGenerator{}
				m.On("GenerateCode", defaultUrlCodeLength).
					Return("abc12345", nil).Once()
				return m
			},
			setupMockRepo: func(ctx context.Context) *mockRepo.Repository {
				m := &mockRepo.Repository{}
				m.On("StoreURLIfNotExists", ctx, "abc12345", "https://test.com", time.Minute).
					Return(false, errors.New("repo error")).Once()
				return m
			},

			inputOriginalURL: "https://test.com",
			inputExpireIn:    1,

			expectedCode:  "",
			expectedError: errors.New("repo error"),
		},
		{
			name: "max retry exceeded",

			setupMockRandomCodeGen: func() *mockRandomCodeGen.CodeGenerator {
				m := &mockRandomCodeGen.CodeGenerator{}
				for i := 0; i < maxRetryAttempts; i++ {
					m.On("GenerateCode", defaultUrlCodeLength).
						Return("dupcode", nil).Once()
				}
				return m
			},
			setupMockRepo: func(ctx context.Context) *mockRepo.Repository {
				m := &mockRepo.Repository{}
				for i := 0; i < maxRetryAttempts; i++ {
					m.On("StoreURLIfNotExists", ctx, "dupcode", "https://test.com", time.Minute).
						Return(false, nil).Once()
				}
				return m
			},

			inputOriginalURL: "https://test.com",
			inputExpireIn:    1,

			expectedCode:  "",
			expectedError: ErrMaxRetryExceeded,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockGen := tc.setupMockRandomCodeGen()
			mockRepo := tc.setupMockRepo(ctx)

			service := NewUrlService(mockRepo, mockGen)

			code, err := service.ShortenURL(ctx, tc.inputOriginalURL, tc.inputExpireIn)

			assert.Equal(t, tc.expectedCode, code)
			assert.Equal(t, tc.expectedError, err)

			mockGen.AssertExpectations(t)
			mockRepo.AssertExpectations(t)
		})
	}
}
