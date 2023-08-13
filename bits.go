package b64

type bits struct {
	size int
	buf  uint16
}

func newBits(size int, i uint16) bits {
	return bits{
		size: size,
		buf:  i,
	}
}

func (b *bits) addLeft(other bits) {
	b.buf = other.buf<<uint16(b.size) | b.buf
	b.size += other.size
}

func (b *bits) cut6SignificantBits() bits {
	if b == nil {
		return bits{}
	}

	remainingBits := b.size - sixBits
	mask := uint16(0b111111) << remainingBits
	maskedBuf := b.buf & mask

	b.size = remainingBits
	b.buf -= maskedBuf

	return newBits(sixBits, maskedBuf>>uint16(remainingBits))
}
