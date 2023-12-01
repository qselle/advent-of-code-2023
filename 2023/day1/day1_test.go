package day1

import "testing"

func TestIsNumberValid(t *testing.T) {
	if isNumber('a') {
		t.Errorf("isNumber('a') = true; want false")
	}
	if !isNumber('0') {
		t.Errorf("isNumber('0') = false; want true")
	}
	if !isNumber('9') {
		t.Errorf("isNumber('9') = false; want true")
	}
}

func TestCharToInt(t *testing.T) {
	if charToInt('0') != 0 {
		t.Errorf("charToInt('0') = %d; want 0", charToInt('0'))
	}
	if charToInt('9') != 9 {
		t.Errorf("charToInt('9') = %d; want 9", charToInt('9'))
	}
}

func TestSolvePartOne(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142
	actual := solvePartOne(input)

	if actual != expected {
		t.Errorf("solvePartOne(%v) = %d; want %d", input, actual, expected)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281
	actual := solvePartTwo(input)

	if actual != expected {
		t.Errorf("solvePartTwo(%v) = %d; want %d", input, actual, expected)
	}
}
