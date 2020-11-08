package util

func removeLeadingAndTrailingZeros(b []byte) ([]byte, error) {
	fnz := getIndexOfFirstNonZeroByte(b)
	lnz := getIndexOfLastNonZeroByte(b)
	if fnz == -1 || lnz == -1 {
		return nil, BytesOperationError("Can not find any non zero bytes")
	}
	rightBoundary := lnz + 1
	return b[fnz:rightBoundary], nil
}

func getIndexOfFirstNonZeroByte(b []byte) int {
	for i, bi := range b {
		if bi != 0 {
			return i
		}
	}
	return -1
}

func getIndexOfLastNonZeroByte(b []byte) int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i
		}
	}
	return -1
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

type BytesOperationError string

func (err BytesOperationError) Error() string {
	return string(err)
}