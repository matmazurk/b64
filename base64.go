package b64

import (
	"fmt"
)

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
)

const (
	sixBits  = 6
	byteSize = 8
)

var b64Map = getB64Map()

func Encode(in []byte) []byte {
	if len(in) == 0 {
		return []byte{}
	}

	res := make([]byte, 0, 2*len(in))
	remaining := newEmptyBits()
	for _, b := range in {
		current := newBits(byteSize, uint16(b))
		current.addLeft(remaining)

		for current.size >= sixBits {
			cut := current.cut6SignificantBits()
			res = append(res, base64Chars[cut.buf])
		}

		remaining = current
	}

	if remaining.size > 0 {
		remaining.addRight(newBits(sixBits-remaining.size, 0))
		res = append(res, base64Chars[remaining.buf])
	}

	for len(res)*sixBits%byteSize != 0 {
		res = append(res, padding)
	}

	return res
}

func Decode(in []byte) ([]byte, error) {
	if len(in) == 0 {
		return []byte{}, nil
	}

	ret := []byte{}
	buf := newEmptyBits()
	appendRetIfBufBigEnough := func() {
		if buf.size >= byteSize {
			cut := buf.cutSignificantBits(byteSize)
			ret = append(ret, byte(cut.buf))
		}
	}
	for _, b := range in {
		appendRetIfBufBigEnough()

		if b == padding {
			return ret, nil
		}

		c, ok := b64Map[b]
		if !ok {
			return nil, fmt.Errorf("invalid base64 character '%c'", b)
		}

		buf.addRight(newBits(sixBits, uint16(c)))
	}

	appendRetIfBufBigEnough()

	return ret, nil
}

func getB64Map() map[byte]int {
	mm := make(map[byte]int, len(base64Chars))
	for i, c := range base64Chars {
		mm[byte(c)] = i
	}
	return mm
}
