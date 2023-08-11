package b64

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
)

const (
	sixBits  = 6
	byteSize = 8
)

type bits struct {
	valid int
	buf   uint16
}

func (b *bits) addLeft(other bits) {
	b.buf = other.buf<<uint16(b.valid) | b.buf
	b.valid += other.valid
}

func (b *bits) cut6SignificantBits() bits {
	if b == nil {
		return bits{}
	}

	remainingBits := b.valid - sixBits
	mask := uint16(0b111111) << remainingBits
	maskedBuf := b.buf & mask

	b.valid = remainingBits
	b.buf -= maskedBuf

	return bits{
		valid: sixBits,
		buf:   maskedBuf >> uint16(remainingBits),
	}
}

func Encode(in []byte) ([]byte, error) {
	if len(in) == 0 {
		return []byte{}, nil
	}

	res := make([]byte, 0, 2*len(in))

	var remaining bits
	for _, b := range in {
		current := bits{
			valid: byteSize,
			buf:   uint16(b),
		}
		current.addLeft(remaining)

		for current.valid >= sixBits {
			cut := current.cut6SignificantBits()
			res = append(res, base64Chars[cut.buf])
		}

		remaining = current
	}

	if remaining.valid > 0 {
		lastByteIndex := remaining.buf << (sixBits - byte(remaining.valid))
		res = append(res, base64Chars[lastByteIndex])
	}

	for len(res)*sixBits%byteSize != 0 {
		res = append(res, padding)
	}

	return res, nil
}
