package main

func getBytesArray(s string) []byte {
	b := []byte(s)
	for i, val := range b {
		b[i] = val - 48
	}
	return b
}

func zeroPadding(s string, maxLen int) string {
	res := s
	for i := len(s); i < maxLen; i++ {
		res = "0" + res
	}
	return res
}

func AddBinary(a string, b string) string {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	aBytes := getBytesArray(zeroPadding(a, maxLen))
	bBytes := getBytesArray(zeroPadding(b, maxLen))

	prevByte := byte(0)
	resBytes := make([]byte, maxLen)
	for i := maxLen - 1; i >= 0; i-- {
		res := aBytes[i] + bBytes[i] + prevByte
		resBytes[i] = res%2 + 48
		prevByte = res / 2
	}

	res := string(resBytes)
	if prevByte == 1 {
		res = "1" + res
	}

	return res
}
