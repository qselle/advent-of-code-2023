package main

import (
	"reflect"
	"testing"
)

func TestParseEngineSchematic(t *testing.T) {
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
	expected := EngineSchematic{
		Parts: []Part{
			{467, []Coordinates{{0, 0}, {0, 1}, {0, 2}}},
			{114, []Coordinates{{0, 5}, {0, 6}, {0, 7}}},
			{35, []Coordinates{{2, 2}, {2, 3}}},
			{633, []Coordinates{{2, 6}, {2, 7}, {2, 8}}},
			{617, []Coordinates{{4, 0}, {4, 1}, {4, 2}}},
			{58, []Coordinates{{5, 7}, {5, 8}}},
			{592, []Coordinates{{6, 2}, {6, 3}, {6, 4}}},
			{755, []Coordinates{{7, 6}, {7, 7}, {7, 8}}},
			{664, []Coordinates{{9, 1}, {9, 2}, {9, 3}}},
			{598, []Coordinates{{9, 5}, {9, 6}, {9, 7}}},
		},
		Symbols: Symbols{
			Coordinates{1, 3}: "*",
			Coordinates{3, 6}: "#",
			Coordinates{4, 3}: "*",
			Coordinates{5, 5}: "+",
			Coordinates{8, 3}: "$",
			Coordinates{8, 5}: "*",
		},
	}
	actual := ParseEngineSchematic(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ParseEngineSchematic(%v) = %v; want %v", input, actual, expected)
	}
}

func TestPartNumbersSum(t *testing.T) {
	input := EngineSchematic{
		Parts: []Part{
			{467, []Coordinates{{0, 0}, {0, 1}, {0, 2}}},
			{114, []Coordinates{{0, 5}, {0, 6}, {0, 7}}},
			{35, []Coordinates{{2, 2}, {2, 3}}},
			{633, []Coordinates{{2, 6}, {2, 7}, {2, 8}}},
			{617, []Coordinates{{4, 0}, {4, 1}, {4, 2}}},
			{58, []Coordinates{{5, 7}, {5, 8}}},
			{592, []Coordinates{{6, 2}, {6, 3}, {6, 4}}},
			{755, []Coordinates{{7, 6}, {7, 7}, {7, 8}}},
			{664, []Coordinates{{9, 1}, {9, 2}, {9, 3}}},
			{598, []Coordinates{{9, 5}, {9, 6}, {9, 7}}},
		},
		Symbols: Symbols{
			Coordinates{1, 3}: "*",
			Coordinates{3, 6}: "#",
			Coordinates{4, 3}: "*",
			Coordinates{5, 5}: "+",
			Coordinates{8, 3}: "$",
			Coordinates{8, 5}: "*",
		},
	}
	expected := 4361
	actual := input.PartNumbersSum()

	if actual != expected {
		t.Errorf("PartNumbersSum(%v) = %d; want %d", input, actual, expected)
	}
}

func TestGearRatiosSum(t *testing.T) {
	input := EngineSchematic{
		Parts: []Part{
			{467, []Coordinates{{0, 0}, {0, 1}, {0, 2}}},
			{114, []Coordinates{{0, 5}, {0, 6}, {0, 7}}},
			{35, []Coordinates{{2, 2}, {2, 3}}},
			{633, []Coordinates{{2, 6}, {2, 7}, {2, 8}}},
			{617, []Coordinates{{4, 0}, {4, 1}, {4, 2}}},
			{58, []Coordinates{{5, 7}, {5, 8}}},
			{592, []Coordinates{{6, 2}, {6, 3}, {6, 4}}},
			{755, []Coordinates{{7, 6}, {7, 7}, {7, 8}}},
			{664, []Coordinates{{9, 1}, {9, 2}, {9, 3}}},
			{598, []Coordinates{{9, 5}, {9, 6}, {9, 7}}},
		},
		Symbols: Symbols{
			Coordinates{1, 3}: "*",
			Coordinates{3, 6}: "#",
			Coordinates{4, 3}: "*",
			Coordinates{5, 5}: "+",
			Coordinates{8, 3}: "$",
			Coordinates{8, 5}: "*",
		},
	}
	expected := 467835
	actual := input.GearRatiosSum()

	if actual != expected {
		t.Errorf("GearRatiosSum(%v) = %d; want %d", input, actual, expected)
	}
}
