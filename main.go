package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"justify/utils"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --align=right something standard")
		return
	}
	if strings.HasPrefix(os.Args[1],"-") {
		if len(os.Args) != 4 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --align=right something standard")
			return
		}

	alignmentFlag := flag.String("align", "left", "output alignment")
	flag.Parse()
	args := flag.Args()

	inputText := args[0]
	banner := args[1]
	alignmentType := *alignmentFlag

	// Determine the banner file and read its content
	bannerFile := utils.DetermineFileName(banner)
	content, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("invalid text file")
		return
	}

	processedText := utils.ReplaceEscape(inputText)
	lines := strings.Split(processedText, "\\n")
	for _, char := range processedText {
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

	justifiedText := utils.AlignJustify(lines, contentLines)
	asciiArtJustified := utils.DisplayText(justifiedText, contentLines)
	asciiArt := utils.DisplayText(inputText, contentLines)

	// Get initial terminal width and print aligned text
	terminalWidth := utils.Getwidth()
	var alignedText string
	if alignmentType == "justify" {
		alignedText = utils.AlignText(asciiArtJustified, terminalWidth, alignmentType)
		fmt.Println(alignedText)
	} else {
		alignedText = utils.AlignText(asciiArt, terminalWidth, alignmentType)
		fmt.Println(alignedText)
	}

	// Listen for terminal resize events and adjust the output
	go func() {
		for {
			newWidth := utils.Getwidth()

			if newWidth != terminalWidth {
				terminalWidth = newWidth
				alignedText = utils.AlignText(asciiArt, terminalWidth, alignmentType)
				fmt.Print("\033[H\033[2J") // Clear screen
				fmt.Println(alignedText)
			}
		}
	}()
	} else {
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . [OPTION] [ALIGN] \n\nEX: go run . hello center")
			return
		}
	
		inputText := os.Args[1]
		banner := "standard"
		alignmentType := os.Args[2]
	
		// Determine the banner file and read its content
		bannerFile := utils.DetermineFileName(banner)
		content, err := os.ReadFile(bannerFile)
		if err != nil {
			fmt.Println("invalid text file")
			return
		}
	
		processedText := utils.ReplaceEscape(inputText)
		lines := strings.Split(processedText, "\\n")
		for _, char := range processedText {
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
	
		justifiedText := utils.AlignJustify(lines, contentLines)
		asciiArtJustified := utils.DisplayText(justifiedText, contentLines)
		asciiArt := utils.DisplayText(inputText, contentLines)
	
		// Get initial terminal width and print aligned text
		terminalWidth := utils.Getwidth()
		var alignedText string
		if alignmentType == "justify" {
			alignedText = utils.AlignText(asciiArtJustified, terminalWidth, alignmentType)
			fmt.Println(alignedText)
		} else {
			alignedText = utils.AlignText(asciiArt, terminalWidth, alignmentType)
			fmt.Println(alignedText)
		}
	
		// Listen for terminal resize events and adjust the output
		go func() {
			for {
				newWidth := utils.Getwidth()
	
				if newWidth != terminalWidth {
					terminalWidth = newWidth
					alignedText = utils.AlignText(asciiArt, terminalWidth, alignmentType)
					fmt.Print("\033[H\033[2J") // Clear screen
					fmt.Println(alignedText)
				}
			}
		}()
	}
}
