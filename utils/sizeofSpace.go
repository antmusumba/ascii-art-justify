package utils

func SizeOfSpace(contentLines []string) int {
	v := ' '
	start := ((v - 32) * 9) + 4
	return len(contentLines[start])
}