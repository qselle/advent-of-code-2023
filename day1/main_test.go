package main

import "testing"

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

func BenchmarkSolvePartOne(b *testing.B) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	for i := 0; i < b.N; i++ {
		solvePartOne(input)
	}
}

func BenchmarkSolvePartTwo(b *testing.B) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	for i := 0; i < b.N; i++ {
		solvePartTwo(input)
	}
}
