package b64

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
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
		actb := uint16(b)
		available := 8
		if rem.validBits > 0 {
			actb = uint16(rem.buf)<<8 | actb
			available += rem.validBits
		}

		for available >= 6 {
			cut := cut6SignificantBits(&actb, available)
			available -= 6
			res = append(res, base64Chars[cut])
		}

		rem = remainder{
			validBits: available,
			buf:       byte(actb),
		}
	}

	if rem.validBits > 0 {
		lastByteIndex := rem.buf << (6 - byte(rem.validBits))
		res = append(res, base64Chars[lastByteIndex])
	}

	for len(res)*6%8 != 0 {
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
