package b64

import "fmt"

type bits struct {
	size int
	buf  uint16
}

func newBits(size int, i uint16) *bits {
	return &bits{
		size: size,
		buf:  i,
	}
}

func (b *bits) equals(other *bits) bool {
	if b == nil || other == nil {
		return b == nil && other == nil
	}

	return *b == *other
}

func (b *bits) addLeft(other *bits) {
	b.buf = other.buf<<uint16(b.size) | b.buf
	b.size += other.size
}

func (b *bits) addRight(other *bits) {
	b.buf = b.buf<<uint16(other.size) | other.buf
	b.size += other.size
}

func (b *bits) cutSignificantBits(n int) *bits {
	if b == nil {
		return &bits{}
	}

	remainingBits := b.size - n
	ones := uint16(pow(2, n+1) - 1)
	mask := ones << remainingBits
	maskedBuf := b.buf & mask

	b.size = remainingBits
	b.buf -= maskedBuf

	return newBits(n, maskedBuf>>uint16(remainingBits))
}

func (b *bits) cut6SignificantBits() *bits {
	return b.cutSignificantBits(sixBits)
}

func (b *bits) String() string {
	return fmt.Sprintf("{size:%d, buf:0b%b}", b.size, b.buf)
}

func pow(x, y int) int {
	res := 1
	for i := 0; i < y-1; i++ {
		res *= x
	}

	return res
}
