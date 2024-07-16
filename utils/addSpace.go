package utils

// AddSpace adds spaces in between words and returns the new string with the added spaces
func AddSpace(word string, space int, contentLines []string) (new string) {
	var sp string
	str := word
	count := 1

	width := Getwidth()
	for GetSizeOfCharacters(str, contentLines)+SizeOfSpace(contentLines) < width {
		str += " "
		count++
	}
	target := count / space

	for len(sp) != target {
		sp += " "
	}

	for _, v := range word {
		if GetSizeOfCharacters(new, contentLines) < width {
			if v == ' ' {
				new += sp
			}
			new += string(v)
		}
	}
	return
}
