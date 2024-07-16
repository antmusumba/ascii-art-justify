package utils

func GetSizeOfCharacters(word string,contentLines []string) int {
	var length int
	for _, v := range word {
		start := ((v - 32) * 9) + 4
		length += len(contentLines[start])
	}
	return length
}