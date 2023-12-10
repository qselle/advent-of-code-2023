package utils

import (
	"os"
	"strconv"
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

// ExtractNumbersFromString convert a string to a slice of int
func ExtractNumbersFromString(str string) []int {
	numbers := make([]int, 0)

	for _, number := range strings.Fields(str) {
		number, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

// ExtractNumberFromStringIgnoringSpaces convert a string to int ignoring spaces
func ExtractNumberFromStringIgnoringSpaces(str string) int {
	number, err := strconv.Atoi(strings.Join(strings.Fields(str), ""))
	if err != nil {
		panic(err)
	}
	return number
}

// GreatestCommonDivisor return the greatest common divisor (https://en.wikipedia.org/wiki/Greatest_common_divisor)
// of two numbers using the Euclidean algorithm (https://en.wikipedia.org/wiki/Euclidean_algorithm see #Implementations)
func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}

// LeastCommonMultiple return the least common multiple (https://en.wikipedia.org/wiki/Least_common_multiple see #Calculation)
// of two numbers
func LeastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)
}
