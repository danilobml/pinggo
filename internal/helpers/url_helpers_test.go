package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid HTTPS URL",
			input:    "https://example.com/path?query=1",
			expected: true,
		},
		{
			name:     "valid FTP URL",
			input:    "ftp://example.com/resource",
			expected: true,
		},
		{
			name:     "missing scheme",
			input:    "example.com",
			expected: false,
		},
		{
			name:     "invalid scheme and characters",
			input:    "ht!tp://bad^url.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := IsValidURL(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
