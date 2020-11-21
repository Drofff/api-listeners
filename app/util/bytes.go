package util

func removeZeros(bytes []byte) []byte {
	nonZerosCount := len(bytes) - countZeros(bytes)
	res := make([]byte, nonZerosCount)
	resIndex := 0
	for _, b := range bytes {
		if b != 0 {
			res[resIndex] = b
			resIndex++
		}
	}
	return res
}

func countZeros(bytes []byte) int {
	c := 0
	for _, b := range bytes {
		if b == 0 {
			c++
		}
 	}
 	return c
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