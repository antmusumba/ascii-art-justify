package utils

func CheckSpace(word string) (check int) {
	for _, v := range word {
		if v == ' ' {
			check++
		}
	}
	return
}