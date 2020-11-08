package util

import "strings"

func ReplaceLine(s string, newLine string, lineNumber int) string {
	lines := strings.Split(s, "\n")
	lines[lineNumber] = newLine
	return strings.Join(lines, "\n")
}
