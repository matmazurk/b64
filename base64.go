package b64

import (
	"fmt"
)

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
)

var b64Map = m()

func m() map[byte]int {
	mm := make(map[byte]int, len(base64Chars))
	for i, c := range base64Chars {
		mm[byte(c)] = i
	}
	return mm
}

const (
	sixBits  = 6
	byteSize = 8
)

func Encode(in []byte) []byte {
	if len(in) == 0 {
		return []byte{}
	}

	res := make([]byte, 0, 2*len(in))

	remaining := newBits(0, 0)
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
		lastByteIndex := remaining.buf << (sixBits - byte(remaining.size))
		res = append(res, base64Chars[lastByteIndex])
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
	buf := newBits(0, 0)
	for _, b := range in {
		if buf.size >= 8 {
			cut := buf.cutSignificantBits(8)
			ret = append(ret, byte(cut.buf))
		}

		if b == padding {
			return ret, nil
		}
		c, ok := b64Map[b]
		if !ok {
			return nil, fmt.Errorf("")
		}

		buf.addRight(newBits(sixBits, uint16(c)))
	}

	if buf.size >= 8 {
		cut := buf.cutSignificantBits(8)
		ret = append(ret, byte(cut.buf))
	}

	return ret, nil
}
