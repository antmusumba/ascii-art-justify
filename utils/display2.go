package utils

import (
	"fmt"
	"os"
	"strings"
)

func DisplayText2(input string, contentLines []string) string{
	if input == "" {
		os.Exit(0)
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		os.Exit(0)
	}
	dat2 := ""
	
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
	dat2 = (strings.Join(linesOfSlice, "\n"))

		}
	}
	return dat2
}