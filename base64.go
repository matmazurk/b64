package b64

const (
	base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padding     = '='
)

func Encode(in []byte) ([]byte, error) {
	if len(in) == 0 {
		return []byte{}, nil
	}

	var res []byte

	var rem struct {
		cnt int
		rem byte
	}
	for _, b := range in {
		actb := uint16(b)
		available := 8
		if rem.cnt > 0 {
			actb = uint16(rem.rem)<<8 | actb
			available += rem.cnt
		}

		for available >= 6 {
			cut := cut6LeftBits(&actb, available)
			available -= 6
			res = append(res, base64Chars[cut])
		}

		rem = struct {
			cnt int
			rem byte
		}{
			cnt: available,
			rem: byte(actb),
		}
	}

	if rem.cnt > 0 {
		aa := rem.rem << (6 - byte(rem.cnt))
		res = append(res, base64Chars[aa])
	}

	for len(res)*6%8 != 0 {
		res = append(res, padding)
	}

	return res, nil
}

func cut6LeftBits(f *uint16, validBytes int) byte {
	if f == nil {
		return 0
	}

	remainingBits := uint16(validBytes - 6)
	mask := uint16(0b111111) << remainingBits

	ret := *f & mask
	*f = *f - ret

	return byte(ret >> remainingBits)
}
