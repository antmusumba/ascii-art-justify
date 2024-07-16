package utils

import "strings"

// Function to align text based on the specified alignment type
func AlignText(text string, width int, alignment string) string {
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

		default:
			lines[i] = line // default to left alignment if type is invalid
		}
	}
	return strings.Join(lines, "\n")
}