package util

func replaceZerosWithSpaceChar(bytes []byte) []byte {
	resBytes := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		if b == 0 {
			resBytes[i] = ' '
		} else {
			resBytes[i] = b
		}
	}
	return resBytes
}

func areBytesEqual(b0 []byte, b1 []byte) bool {
	if len(b0) != len(b1) {
		return false
	}
	for i := 0; i < len(b0); i++ {
		if b0[i] != b1[i] {
			return false
		}
	}
	return true
}