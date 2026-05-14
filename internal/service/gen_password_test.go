package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		expectedLength int
		expectedErr    string // dùng string thay vì error để compare chuẩn
	}{
		{name: "valid length 8", expectedLength: 8, expectedErr: ""},
		{name: "valid length 16", expectedLength: 16, expectedErr: ""},
		{name: "valid length 32", expectedLength: 32, expectedErr: ""},
		{name: "length zero", expectedLength: 0, expectedErr: "length must be greater than 0"},
		{name: "negative length", expectedLength: -1, expectedErr: "length must be greater than 0"},
	}

	testSvc := NewGenPassword()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := testSvc.GeneratePassword(tc.expectedLength)

			// ===== Error check =====
			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err.Error())
				return
			}

			// ===== Success cases =====
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedLength, len(result))

			// optional: check charset validity
			for i := 0; i < len(result); i++ {
				assert.Contains(t, defaultCharset, string(result[i]))
			}
		})
	}
}
