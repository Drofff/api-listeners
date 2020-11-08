package util

import "testing"

func TestRemoveLeadingAndTrailingZeros(t *testing.T) {
	testBytes := []byte{ 0, 0, 4, 5, 1, 2, 63, 0, 0, 0 }
	expectedResult := []byte{ 4, 5, 1, 2, 63 }

	actualResult, err := removeLeadingAndTrailingZeros(testBytes)
	if err != nil {
		t.Error(err)
	}

	if !areBytesEqual(expectedResult, actualResult) {
		t.Error("Unexpected result:", actualResult)
	}
}