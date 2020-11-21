package util

import "testing"

func TestRemoveZeros(t *testing.T) {
	testBytes := []byte{ 0, 0, 3, 1, 2, 0, 4, 0, 0, 1, 0 }
	expectedResult := []byte{ 3, 1, 2, 4, 1 }
	actualResult := removeZeros(testBytes)
	if !areBytesEqual(expectedResult, actualResult) {
		t.Errorf("Incorrectly removed zero bytes. Expected: %v, returned: %v", expectedResult, actualResult)
	}
}