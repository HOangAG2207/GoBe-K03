package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GeneratePassword(t *testing.T) {

	// Enable parallel execution for faster test runtime
	t.Parallel()

	// =========================
	// 1. Test cases definition
	// =========================
	testCases := []struct {
		name           string // test case name (display only)
		expectedLength int    // input password length
		expectedErr    string // expected error message (empty = no error expected)
	}{
		// ===== valid cases =====
		{name: "valid length 8", expectedLength: 8, expectedErr: ""},
		{name: "valid length 16", expectedLength: 16, expectedErr: ""},
		{name: "valid length 32", expectedLength: 32, expectedErr: ""},

		// ===== invalid cases =====
		{name: "length zero", expectedLength: 0, expectedErr: "length must be greater than 0"},
		{name: "negative length", expectedLength: -1, expectedErr: "length must be greater than 0"},
	}

	// Initialize service instance once for all test cases
	testSvc := NewGenPassword()

	// =========================
	// 2. Execute test cases
	// =========================
	for _, tc := range testCases {

		// Capture range variable to avoid parallel issue
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			// Run subtests in parallel
			t.Parallel()

			// =========================
			// 3. Call service method
			// =========================
			result, err := testSvc.GeneratePassword(tc.expectedLength)

			// =========================
			// 4. Validate error cases
			// =========================
			if tc.expectedErr != "" {

				// Expect an error
				assert.Error(t, err)

				// Compare exact error message
				assert.Equal(t, tc.expectedErr, err.Error())

				return
			}

			// =========================
			// 5. Validate success cases
			// =========================
			assert.NoError(t, err)

			// Ensure password length matches expectation
			assert.Equal(t, tc.expectedLength, len(result))

			// =========================
			// 6. Validate charset correctness
			// =========================
			// Ensure every generated character belongs to allowed charset
			for i := 0; i < len(result); i++ {
				assert.Contains(t, defaultCharset, string(result[i]))
			}
		})
	}
}
