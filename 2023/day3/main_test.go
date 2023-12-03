package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 4361
	actual := solvePartOne(input)

	if actual != expected {
		t.Errorf("solvePartOne(%v) = %d; want %d", input, actual, expected)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 467835
	actual := solvePartTwo(input)

	if actual != expected {
		t.Errorf("solvePartTwo(%v) = %d; want %d", input, actual, expected)
	}
}
