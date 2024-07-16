package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"ascii-art/utils"
)

// Struct to store terminal window size
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Function to get terminal size
func getTerminalSize() (int, int, error) {
	ws := &winsize{}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if errno != 0 {
		return 0, 0, errno
	}
	return int(ws.Col), int(ws.Row), nil
}

// Function to align text based on the specified alignment type
func alignText(text string, width int, alignment string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if len(line) > width {
			line = line[:width]
		}
		switch alignment {
		case "left":
			lines[i] = line
		case "right":
			spaces := width - len(line)
			if spaces > 0 {
				lines[i] = strings.Repeat(" ", spaces) + line
			}
		case "center":
			spaces := (width - len(line)) / 2
			if spaces > 0 {
				lines[i] = strings.Repeat(" ", spaces) + line
			}
		case "justify":
			// For justify, we will handle it separately
		default:
			lines[i] = line // default to left alignment if type is invalid
		}
	}
	return strings.Join(lines, "\n")
}

// Function to justify align text
func alignJustify(words []string, contentLines []string) string {
	var justifiedLines []string
	for _, word := range words {
		if word == "" {
			justifiedLines = append(justifiedLines, word)
			continue
		} else {
			spaces := utils.CheckSpace(word)
			if spaces != 0 {
				word = utils.AddSpace(word, spaces, contentLines)
			}
			justifiedLines = append(justifiedLines, word)
		}
	}
	words = justifiedLines
	return strings.Join(words, " ")
}

func main() {
	alignmentFlag := flag.String("align", "left", "output alignment")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --align=right something standard")
		return
	}

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

	justifiedText := alignJustify(lines, contentLines)
	asciiArtJustified := utils.DisplayText(justifiedText, contentLines)
	asciiArt := utils.DisplayText(inputText, contentLines)

	// Get initial terminal width and print aligned text
	terminalWidth, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}
	var alignedText string
	if alignmentType == "justify" {
		alignedText = alignText(asciiArtJustified, terminalWidth, alignmentType)
		fmt.Println(alignedText)
	} else {
		alignedText = alignText(asciiArt, terminalWidth, alignmentType)
		fmt.Println(alignedText)
	}

	// Listen for terminal resize events and adjust the output
	go func() {
		for {
			newWidth, _, err := getTerminalSize()
			if err != nil {
				fmt.Println("Error getting terminal size:", err)
				continue
			}
			if newWidth != terminalWidth {
				terminalWidth = newWidth
				alignedText = alignText(asciiArt, terminalWidth, alignmentType)
				fmt.Print("\033[H\033[2J") // Clear screen
				fmt.Println(alignedText)
			}
		}
	}()
}
