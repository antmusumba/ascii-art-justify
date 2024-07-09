package utils

import "strings"

func ReplaceEscape(input string) string {
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	input = strings.ReplaceAll(input, "\\r", "\r")
	input = strings.ReplaceAll(input, "\\v", "\v")
	input = strings.ReplaceAll(input, "\\f", "\f")
	input = strings.ReplaceAll(input, "\\a", "\a")

	return input
}
