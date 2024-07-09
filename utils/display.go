package utils

import (
	"fmt"
	"strings"
)

func DisplayText(align string,input string, contentLines []string) {
	if input == "" {
		return
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}
	
	// split the input string with the "\\n" into a slice strings
	wordslice := strings.Split(input, "\\n")

	count := 0
	for _, word := range wordslice {
		if word == "" {
			count++
			if count < len(wordslice) {
				fmt.Println()
			}
		} else {
			PrintWord(word, contentLines)
		}
	}
}
