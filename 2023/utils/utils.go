package utils

import (
	"os"
	"strings"
)

// ReadFileByLine open and store a file by line in an array
// of string
func ReadFileByLine(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// RuneToInt convert a rune to int
func RuneToInt(char rune) int {
	return int(char - '0')
}
