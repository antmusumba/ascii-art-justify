package utils

import (
	"fmt"
	"strings"
)

func DisplayText2(input string, contentLines []string) {
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
			
	linesOfSlice := make([]string, 9)
	for _, v := range word {
		for i := 1; i <= 9; i++ {
			linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
		}
	}
	fmt.Print(strings.Join(linesOfSlice, "\n"))

		}
	}
}