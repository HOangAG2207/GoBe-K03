package test_endpoint

import (
	"net/http/httptest"
	"testing"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestEndpoint_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		setupTestHttp        func(e config.Engine) *httptest.ResponseRecorder
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "success generate password",
			setupTestHttp: func(e config.Engine) *httptest.ResponseRecorder {

				// ===== create request =====
				req := httptest.NewRequest("GET", "/api/gen-password", nil)
				rec := httptest.NewRecorder()

				// ===== execute through real HTTP server =====
				e.ServeHTTP(rec, req)

				return rec
			},
			expectedStatusCode:   200,
			expectedResponseBody: "password",
		},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// ===== init real engine (full stack) =====
			cfg := config.Load()
			engine := config.NewEngine(&config.EngineOpts{
				Cfg: cfg,
			})

			// ===== execute request =====
			rec := tc.setupTestHttp(engine)

			// ===== assert status code =====
			assert.Equal(t, tc.expectedStatusCode, rec.Code)

			// ===== assert response body contains expected string =====
			assert.Contains(t, rec.Body.String(), tc.expectedResponseBody)
		})
	}
}
