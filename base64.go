package b64

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
)

const (
	sixBits  = 6
	byteSize = 8
)

func Encode(in []byte) ([]byte, error) {
	if len(in) == 0 {
		return []byte{}, nil
	}

	res := make([]byte, 0, 2*len(in))

	var remaining bits
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

	return res, nil
}
