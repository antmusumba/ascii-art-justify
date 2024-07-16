package utils

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)


func Getwidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	sizeArr := strings.Fields(string(out))
	width, _ := strconv.Atoi(sizeArr[1])
	return width
}