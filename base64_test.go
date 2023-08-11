package b64_test

import (
	"testing"

	"github.com/matmazurk/b64"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	emptyBytes := []byte{}
	tcs := []struct {
		in       []byte
		expected []byte
	}{
		{
			in:       emptyBytes,
			expected: emptyBytes,
		},
		{
			in:       []byte("Many hands make light work."),
			expected: []byte("TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"),
		},
		{
			in:       []byte("M"),
			expected: []byte("TQ=="),
		},
		{
			in:       []byte("Ma"),
			expected: []byte("TWE="),
		},
	}

	for _, tc := range tcs {
		res, err := b64.Encode(tc.in)
		require.NoError(t, err)

		require.Equal(t, tc.expected, res)
	}
}
