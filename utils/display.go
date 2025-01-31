package utils

import (
	"fmt"
	"strings"
)

func DisplayText(input string, contentLines []string) string {
	data := ""
	// split the input string with the "\\n" into a slice strings
	wordslice := strings.Split(input, "\\n")

	count := 1
	
	for _, word := range wordslice {
		if word == "" {
			count++
			if count < len(wordslice) {
				fmt.Println()
			}
		} else {
			data += PrintWord(word, contentLines)
		}
	}
	return data
}
