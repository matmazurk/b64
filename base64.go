package b64

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='

	base64Size = 6
	byteSize   = 8
)

func Encode(in []byte) ([]byte, error) {
	if len(in) == 0 {
		return []byte{}, nil
	}

	res := make([]byte, 0, 2*len(in))

	type remainder struct {
		validBits int
		buf       byte
	}

	var rem remainder
	for _, b := range in {
		currByte := uint16(b)
		validBits := byteSize
		if rem.validBits > 0 {
			currByte = uint16(rem.buf)<<byteSize | currByte
			validBits += rem.validBits
		}

		for validBits >= base64Size {
			cut := cut6SignificantBits(&currByte, validBits)
			validBits -= base64Size
			res = append(res, base64Chars[cut])
		}

		rem = remainder{
			validBits: validBits,
			buf:       byte(currByte),
		}
	}

	if rem.validBits > 0 {
		lastByteIndex := rem.buf << (base64Size - byte(rem.validBits))
		res = append(res, base64Chars[lastByteIndex])
	}

	for len(res)*base64Size%byteSize != 0 {
		res = append(res, padding)
	}

	return res, nil
}

func cut6SignificantBits(v *uint16, validBytes int) byte {
	if v == nil {
		return 0
	}

	remainingBits := uint16(validBytes - 6)
	mask := uint16(0b111111) << remainingBits

	ret := *v & mask
	*v = *v - ret

	return byte(ret >> remainingBits)
}
