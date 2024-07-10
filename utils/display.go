package utils

import (
	"fmt"
	"os"
	"strings"
)

func DisplayText(input string, contentLines []string) string {
	if input == "" {
		os.Exit(0)
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		os.Exit(0)
	}
	data := ""
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
			data = PrintWord(word, contentLines)
		}
	}
	return data
}
