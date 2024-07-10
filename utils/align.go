package utils

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

// getTerminalSize returns the width and height of the terminal.
func getTerminalSize() (int, int, error) {
	ws := &struct {
		Row, Col, Xpixel, Ypixel uint16
	}{}
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

// printAdjustedOutput adjusts the output to fit the terminal width.
func printAdjustedOutput(str string) {
	width, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}


	// Split the text into lines based on the terminal width
	words := strings.Split(str, " ")
	var line string
	for _, word := range words {
		if len(line)+len(word)+1 > width {
			fmt.Println(line)
			line = word
		} else {
			if len(line) > 0 {
				line += " "
			}
			line += word
		}
	}
	if len(line) > 0 {
		fmt.Println(line)
	}
}


