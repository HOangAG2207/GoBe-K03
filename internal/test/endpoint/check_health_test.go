package test_endpoint

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/HOangAG2207/GoBe-K03/internal/config"
	"github.com/stretchr/testify/assert"
)

// CheckHealthResponse represents the expected JSON response structure
type CheckHealthResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

// TestEndpoint_CheckHealth performs an end-to-end test for the health-check endpoint.
// It verifies the full HTTP flow: routing → handler → service → response.
func TestEndpoint_CheckHealth(t *testing.T) {
	t.Parallel()

	// Table-driven test cases for endpoint testing
	testCases := []struct {
		name string

		// setupTestHTTP executes a real HTTP request against the server
		setupTestHTTP func(server config.Engine) *httptest.ResponseRecorder

		expectedStatusCode int
		expectedMessage    string
		expectedService    string
	}{
		{
			name: "success check health",

			// Create and execute HTTP request through real server stack
			setupTestHTTP: func(server config.Engine) *httptest.ResponseRecorder {

				// Create HTTP request
				req := httptest.NewRequest("GET", "/api/health-check", nil)
				rec := httptest.NewRecorder()

				// Send request through real HTTP engine
				server.ServeHTTP(rec, req)

				return rec
			},

			// Expected HTTP response status
			expectedStatusCode: 200,

			// Expected response content
			expectedMessage: "OK",
			expectedService: "GoBe-K03",
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Initialize full application engine (integration test level)
			cfg := config.Load()
			engine := config.NewEngine(&config.EngineOpts{
				Cfg: cfg,
			})

			// Execute test request
			rec := tc.setupTestHTTP(engine)

			// Assert HTTP status code
			assert.Equal(t, tc.expectedStatusCode, rec.Code)

			// Decode JSON response
			var res CheckHealthResponse
			err := json.Unmarshal(rec.Body.Bytes(), &res)
			assert.NoError(t, err)

			// Assert response fields
			assert.Equal(t, tc.expectedMessage, res.Message)
			assert.Equal(t, tc.expectedService, res.ServiceName)

			// Instance ID should always be generated and non-empty
			assert.NotEmpty(t, res.InstanceID)
		})
	}
}
