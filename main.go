package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art/utils"
)

func main() {
	// Check if the number of arguments is less than 2 or greater than 5
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if strings.HasPrefix(os.Args[1], "-") {
		if len(os.Args) != 4 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}

		AlignVar := flag.String("align", "left", "output alignment")
		flag.Parse()
		arg := flag.Args()

		inputWord := arg[0] // The string to be converted to ASCII art
		banner := arg[1]    // The banner style to use
		alignType := *AlignVar

		// Load ASCII characters from file
		file := utils.DetermineFileName(banner)
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("invalid text file")
			return
		}
		s := utils.ReplaceEscape(os.Args[1])
		// Check for in	fmt.Println("the count is ",count)valid characters
		for _, char := range s {
			if char > 126 || char < 32 {
				fmt.Printf("Error: Character %q is not accepted\n", char)
				os.Exit(0)
			}
		}

		contentLines := utils.SplitFile(string(content))

		if len(contentLines) != 856 {
			fmt.Println("invalid text file")
			return
		}
		utils.DisplayText(alignType, inputWord, contentLines)
	} else {
		var file string
		if len(os.Args) == 2 {
			file = "standard.txt"
		} else if len(os.Args) == 3 {
			file = utils.DetermineFileName(os.Args[2])
		}
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("invalid text file")
			return
		}
		s := utils.ReplaceEscape(os.Args[1])
		// Check for in	fmt.Println("the count is ",count)valid characters
		for _, char := range s {
			if char > 126 || char < 32 {
				fmt.Printf("Error: Character %q is not accepted\n", char)
				os.Exit(0)
			}
		}

		contentLines := utils.SplitFile(string(content))

		if len(contentLines) != 856 {
			fmt.Println("invalid text file")
			return
		}
		utils.DisplayText2(os.Args[1], contentLines)
	}
}
