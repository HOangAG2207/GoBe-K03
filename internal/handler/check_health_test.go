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

func TestHandler_CheckHealth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest     func(ctx echo.Context)
		setupMockService func(t *testing.T, ctx echo.Context) *mocks.HealthCheck

		expectedStatus   int
		expectedResponse service.HealthCheckResponse
	}{
		{
			name: "should return OK response",
			setupRequest: func(ctx echo.Context) {
				req := httptest.NewRequest(http.MethodGet, "/health-check", bytes.NewBuffer(nil))
				rec := httptest.NewRecorder()
				ctx.SetRequest(req)
				ctx.SetResponse(echo.New().NewContext(req, rec).Response())
			},
			setupMockService: func(t *testing.T, ctx echo.Context) *mocks.HealthCheck {
				mockSvc := new(mocks.HealthCheck)

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
			expectedStatus: http.StatusOK,
			expectedResponse: service.HealthCheckResponse{
				Message:     "OK",
				ServiceName: "test-service",
				InstanceID:  "instance-123",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			mockSvc := tc.setupMockService(t, ctx)
			h := NewHealthCheck(mockSvc)

			err := h.CheckHealth(ctx)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			mockSvc.AssertExpectations(t)
		})
	}
}
