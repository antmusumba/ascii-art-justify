package utils

import "strings"


// Function to justify align text
func AlignJustify(words []string, contentLines []string) string {
	var justifiedLines []string
	for _, word := range words {
		if word == "" {
			justifiedLines = append(justifiedLines, word)
			continue
		} else {
			spaces := CheckSpace(word)
			if spaces != 0 {
				word = AddSpace(word, spaces, contentLines)
			}
			justifiedLines = append(justifiedLines, word)
		}
	}
	words = justifiedLines
	return strings.Join(words, " ")
}