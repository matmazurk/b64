package b64_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matmazurk/b64"
	"github.com/stretchr/testify/require"
)

func TestEncodeAndDecode(t *testing.T) {
	emptyBytes := []byte{}
	tcs := []struct {
		in  []byte
		out []byte
	}{
		{
			in:  emptyBytes,
			out: emptyBytes,
		},
		{
			in:  []byte("Many hands make light work."),
			out: []byte("TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"),
		},
		{
			in:  []byte("M"),
			out: []byte("TQ=="),
		},
		{
			in:  []byte("Ma"),
			out: []byte("TWE="),
		},
	}

	for _, tc := range tcs {
		t.Run(
			fmt.Sprintf("%s encoding should result in %s", string(tc.in), string(tc.out)),
			func(t *testing.T) {
				res := b64.Encode(tc.in)
				require.Equal(t, tc.out, res)
			},
		)
	}

	for _, tc := range tcs {
		t.Run(
			fmt.Sprintf("%s decoding should result in %s", string(tc.out), string(tc.in)),
			func(t *testing.T) {
				res, err := b64.Decode(tc.out)
				require.NoError(t, err)
				require.Equal(t, tc.in, res)
			},
		)
	}

	t.Run("Decode should fail for input with invalid base64 character", func(t *testing.T) {
		const input = "∂"
		_, err := b64.Decode([]byte(input))
		require.ErrorContains(t, err, "invalid base64 character")
	})
}

func FuzzEncodeAndDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, bt []byte) {
		enc := b64.Encode(bt)
		dec, err := b64.Decode(enc)
		require.NoError(t, err)
		if !bytes.Equal(bt, dec) {
			t.Errorf("decoded different than original input: %s", cmp.Diff(bt, dec))
		}
	})
}
