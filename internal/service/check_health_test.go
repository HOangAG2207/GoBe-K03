package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CheckHealth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		inputServiceName    string
		inputInstanceID     string
		expectedMessage     string
		expectedServiceName string
		expectedInstanceID  string
		expectUUIDGenerated bool
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

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testSvc := NewHealthCheck(tc.inputServiceName, tc.inputInstanceID)
			resp := testSvc.CheckHealth()

			assert.Equal(t, tc.expectedMessage, resp.Message)
			assert.Equal(t, tc.expectedServiceName, resp.ServiceName)

			if tc.expectUUIDGenerated {
				assert.NotEmpty(t, resp.InstanceID)
			} else {
				assert.Equal(t, tc.expectedInstanceID, resp.InstanceID)
			}
		})
	}
}
