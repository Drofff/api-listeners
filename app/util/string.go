package util

import (
	"io"
	"strings"
)

func ReplaceLine(s string, newLine string, lineNumber int) string {
	lines := strings.Split(s, "\n")
	lines[lineNumber] = newLine
	return strings.Join(lines, "\n")
}

func IsBlank(s string) bool {
	sTrimmed := strings.TrimSpace(s)
	return len(sTrimmed) == 0
}

func ReadAsString(r io.Reader) (string, error) {
	sb := new(strings.Builder)
	_, err := io.Copy(sb, r)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}