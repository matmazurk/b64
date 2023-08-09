package b64_test

import (
	"testing"

	"github.com/matmazurk/b64"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	tcs := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "for empty string should return empty string",
			in:       "",
			expected: "",
		},
		{
			name:     "should correctly encode first example",
			in:       "Many hands make light work.",
			expected: "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := b64.Encode(tc.in)
			require.NoError(t, err)

			require.Equal(t, tc.expected, res)
		})
	}
}
