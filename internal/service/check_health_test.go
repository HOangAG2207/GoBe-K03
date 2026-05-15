package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestService_CheckHealth verifies the behavior of the CheckHealth service method.
// It ensures correct response formatting and instance ID handling logic.
func TestService_CheckHealth(t *testing.T) {
	t.Parallel()

	// Table-driven test cases for different input scenarios
	testCases := []struct {
		name                string // test case name
		inputServiceName    string // service name passed to constructor
		inputInstanceID     string // instance ID passed to constructor
		expectedMessage     string // expected response message
		expectedServiceName string // expected service name in response
		expectedInstanceID  string // expected instance ID (if not generated)
		expectUUIDGenerated bool   // flag to check if UUID should be auto-generated
	}{
		{
			name:                "should return given instanceID",
			inputServiceName:    "test-service",
			inputInstanceID:     "instance-123",
			expectedMessage:     "OK",
			expectedServiceName: "test-service",
			expectedInstanceID:  "instance-123",
			expectUUIDGenerated: false,
		},
		{
			name:                "should generate instanceID when empty",
			inputServiceName:    "test-service",
			inputInstanceID:     "",
			expectedMessage:     "OK",
			expectedServiceName: "test-service",
			expectedInstanceID:  "",
			expectUUIDGenerated: true,
		},
	}

	// Iterate through all test cases
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Initialize service with test inputs
			testSvc := NewHealthCheck(tc.inputServiceName, tc.inputInstanceID)

			// Call the business logic
			resp := testSvc.CheckHealth()

			// Assert response message is always "OK"
			assert.Equal(t, tc.expectedMessage, resp.Message)

			// Assert service name is correctly propagated
			assert.Equal(t, tc.expectedServiceName, resp.ServiceName)

			// Assert instance ID behavior:
			// - If UUID is expected, it should not be empty
			// - Otherwise, it should match the input value
			if tc.expectUUIDGenerated {
				assert.NotEmpty(t, resp.InstanceID)
			} else {
				assert.Equal(t, tc.expectedInstanceID, resp.InstanceID)
			}
		})
	}
}
