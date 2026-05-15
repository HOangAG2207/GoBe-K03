package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HOangAG2207/GoBe-K03/internal/service"
	"github.com/HOangAG2207/GoBe-K03/internal/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestHandler_CheckHealth tests the CheckHealth handler
// to ensure it returns the correct response and properly calls the service layer.
func TestHandler_CheckHealth(t *testing.T) {
	t.Parallel()

	// Table-driven test cases
	testCases := []struct {
		name string

		// setupRequest prepares the HTTP request and Echo context
		setupRequest func(ctx echo.Context)

		// setupMockService initializes the mock service and defines expected behavior
		setupMockService func(t *testing.T, ctx echo.Context) *mocks.HealthCheck

		// expected HTTP status code
		expectedStatus int

		// expected response from the service layer (business expectation)
		expectedResponse service.HealthCheckResponse
	}{
		{
			name: "should return OK response",

			// Create a fake GET request for /health-check
			setupRequest: func(ctx echo.Context) {
				req := httptest.NewRequest(http.MethodGet, "/health-check", bytes.NewBuffer(nil))
				rec := httptest.NewRecorder()

				// Attach request to Echo context
				ctx.SetRequest(req)

				// Attach response writer to Echo context
				ctx.SetResponse(echo.New().NewContext(req, rec).Response())
			},

			// Mock the service layer to isolate handler testing
			setupMockService: func(t *testing.T, ctx echo.Context) *mocks.HealthCheck {
				mockSvc := new(mocks.HealthCheck)

				// Expect CheckHealth to be called exactly once
				// and return a predefined response
				mockSvc.
					On("CheckHealth").
					Return(service.HealthCheckResponse{
						Message:     "OK",
						ServiceName: "test-service",
						InstanceID:  "instance-123",
					}).
					Once()

				return mockSvc
			},

			// Expected HTTP status code
			expectedStatus: http.StatusOK,

			// Expected business response (not directly asserted in this test yet)
			expectedResponse: service.HealthCheckResponse{
				Message:     "OK",
				ServiceName: "test-service",
				InstanceID:  "instance-123",
			},
		},
	}

	// Iterate through all test cases
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a new Echo instance per test case to avoid shared state issues
			e := echo.New()

			// Create HTTP request and response recorder
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()

			// Create Echo context
			ctx := e.NewContext(req, rec)

			// Initialize mock service
			mockSvc := tc.setupMockService(t, ctx)

			// Create handler with dependency injection
			h := NewHealthCheck(mockSvc)

			// Call the handler function
			err := h.CheckHealth(ctx)

			// Assert no error occurred
			assert.NoError(t, err)

			// Assert HTTP status code
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Verify that all expected mock calls were executed
			mockSvc.AssertExpectations(t)
		})
	}
}
