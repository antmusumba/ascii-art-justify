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

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

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
			words := strings.Fields(line)
			if len(words) < 2 {
				lines[i] = line
				continue
			}
			spacesNeeded := width - len(line) + len(words) - 1
			spaceWidth := spacesNeeded / (len(words) - 1)
			extraSpaces := spacesNeeded % (len(words) - 1)
			var justifiedLine string
			for j, word := range words {
				if j > 0 {
					spaceToAdd := spaceWidth
					if j <= extraSpaces {
						spaceToAdd++
					}
					justifiedLine += strings.Repeat(" ", spaceToAdd)
				}
				justifiedLine += word
			}
			lines[i] = justifiedLine
		default:
			lines[i] = line // default to left alignment if type is invalid
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	alignVar := flag.String("align", "left", "output alignment")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	inputWord := args[0]
	banner := args[1]
	alignType := *alignVar

	file := utils.DetermineFileName(banner)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("invalid text file")
		return
	}
	s := utils.ReplaceEscape(inputWord)

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

	data1 := utils.DisplayText(inputWord, contentLines)

	// Print initially aligned text
	width, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}
	alignedText := alignText(data1, width, alignType)
	fmt.Println(alignedText)

	// Listen for terminal resize events and adjust the output
	go func() {
		for {
			// Check terminal size
			newWidth, _, err := getTerminalSize()
			if err != nil {
				fmt.Println("Error getting terminal size:", err)
				continue
			}
			if newWidth != width {
				width = newWidth
				// Re-align text and print
				alignedText = alignText(data1, width, alignType)
				fmt.Print("\033[H\033[2J") // Clear screen
				fmt.Println(alignedText)
			}
		}
	}()

	// Wait for user input to exit
	fmt.Println("Press 'Enter' to exit...")
	var input string
	fmt.Scanln(&input)
}

