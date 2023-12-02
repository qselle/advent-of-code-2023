package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	digitByLetters = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
)

func readFileByLine(inputFile string) []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func isNumber(char uint8) bool {
	return char >= '0' && char <= '9'
}

func charToInt(char uint8) int {
	return int(char - '0')
}

func solvePartTwo(input []string) (total int) {
	wg := sync.WaitGroup{}

	for _, line := range input {
		firstDigit, lastDigit := 0, 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := 0; index < len(line); index++ {
				char := line[index]
				if isNumber(char) {
					firstDigit = charToInt(char)
					return
				}

				for digit, digitByLetter := range digitByLetters {
					if strings.HasPrefix(line[index:], digitByLetter) {
						firstDigit = digit + 1
						return
					}
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := len(line) - 1; index >= 0; index-- {
				char := line[index]
				if isNumber(char) {
					lastDigit = charToInt(char)
					return
				}
				for digit, digitByLetter := range digitByLetters {
					if strings.HasSuffix(line[:index+1], digitByLetter) {
						lastDigit = digit + 1
						return
					}
				}
			}
		}()

		wg.Wait()
		total += firstDigit*10 + lastDigit
	}
	return

}

func solvePartOne(input []string) (total int) {
	wg := sync.WaitGroup{}

	for _, line := range input {
		firstDigit, lastDigit := 0, 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := 0; index < len(line); index++ {
				char := line[index]
				if isNumber(char) {
					firstDigit = charToInt(char)
					return
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := len(line) - 1; index >= 0; index-- {
				char := line[index]
				if isNumber(char) {
					lastDigit = charToInt(char)
					return
				}
			}
		}()

		wg.Wait()
		total += firstDigit*10 + lastDigit
	}
	return
}

func main() {
	input := readFileByLine("input.txt")
	fmt.Println("Part 1:", solvePartOne(input))
	fmt.Println("Part 2:", solvePartTwo(input))
}
