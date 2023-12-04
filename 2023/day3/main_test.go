package main

import "testing"

func TestPartNumbersSum(t *testing.T) {
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
	es := ParseEngineSchematic(input)
	expected := 4361
	actual := es.PartNumbersSum()

	if actual != expected {
		t.Errorf("PartNumbersSum(%v) = %d; want %d", input, actual, expected)
	}
}

func TestGearRatiosSum(t *testing.T) {
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
	es := ParseEngineSchematic(input)
	expected := 467835
	actual := es.GearRatiosSum()

	if actual != expected {
		t.Errorf("GearRatiosSum(%v) = %d; want %d", input, actual, expected)
	}
}
