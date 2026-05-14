package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HOangAG2207/GoBe-K03/internal/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GeneratePassword(t *testing.T) {
	// Enable parallel execution for faster test run
	t.Parallel()

	// =========================
	// 1. Define test cases
	// =========================
	testCases := []struct {
		name               string // test case name
		queryParam         string // query param "length"
		mockReturnPassword string // mocked password from service
		mockReturnError    error  // mocked error from service
		expectedStatusCode int    // expected HTTP status code
		expectedContains   string // expected response body content
		expectServiceCall  bool   // whether service should be called
	}{
		// ===== success cases =====
		{
			name:               "success default length",
			queryParam:         "",
			mockReturnPassword: "abc123",
			mockReturnError:    nil,
			expectedStatusCode: http.StatusOK,
			expectedContains:   "abc123",
			expectServiceCall:  true,
		},
		{
			name:               "success custom length",
			queryParam:         "16",
			mockReturnPassword: "xyz999",
			mockReturnError:    nil,
			expectedStatusCode: http.StatusOK,
			expectedContains:   "xyz999",
			expectServiceCall:  true,
		},

		// ===== validation failure cases =====
		{
			name:               "invalid length",
			queryParam:         "-1",
			mockReturnPassword: "",
			mockReturnError:    nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedContains:   "invalid length",
			expectServiceCall:  false,
		},

		// ===== service failure cases =====
		{
			name:               "service error",
			queryParam:         "10",
			mockReturnPassword: "",
			mockReturnError:    errors.New("service failed"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedContains:   "Internal Server Error",
			expectServiceCall:  true,
		},
	}

	// =========================
	// 2. Execute test cases
	// =========================
	for _, tc := range testCases {

		// capture variable to avoid race condition in parallel tests
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			// run subtests in parallel
			t.Parallel()

			// =========================
			// 3. Setup Echo context
			// =========================
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/gen-pass", nil)

			// inject query parameter if provided
			if tc.queryParam != "" {
				q := req.URL.Query()
				q.Add("length", tc.queryParam)
				req.URL.RawQuery = q.Encode()
			}

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// =========================
			// 4. Setup mock service
			// =========================
			mockSvc := new(mocks.GenPassword)

			// only expect service call when handler should call service
			if tc.expectServiceCall {
				mockSvc.
					On("GeneratePassword", mock.Anything).
					Return(tc.mockReturnPassword, tc.mockReturnError)
			}

			// =========================
			// 5. Initialize handler
			// =========================
			h := NewGenPassword(mockSvc)

			// =========================
			// 6. Execute handler
			// =========================
			err := h.GeneratePassword(c)

			// handler should not return error directly
			assert.NoError(t, err)

			// =========================
			// 7. Assert HTTP response
			// =========================
			assert.Equal(t, tc.expectedStatusCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tc.expectedContains)

			// =========================
			// 8. Assert mock expectations
			// =========================
			// only validate when service is expected to be called
			if tc.expectServiceCall {
				mockSvc.AssertExpectations(t)
			}
		})
	}
}
