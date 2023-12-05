package main

import (
	"reflect"
	"testing"
)

func TestParsePuzzleInput(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := []Game{
		{
			id: 1,
			sets: []Set{
				{red: 4, blue: 3, green: 0},
				{red: 1, blue: 6, green: 2},
				{red: 0, blue: 0, green: 2},
			},
		},
		{
			id: 2,
			sets: []Set{
				{red: 0, blue: 1, green: 2},
				{red: 1, blue: 4, green: 3},
				{red: 0, blue: 1, green: 1},
			},
		},
		{
			id: 3,
			sets: []Set{
				{red: 20, blue: 6, green: 8},
				{red: 4, blue: 5, green: 13},
				{red: 1, blue: 0, green: 5},
			},
		},
		{
			id: 4,
			sets: []Set{
				{red: 3, blue: 6, green: 1},
				{red: 6, blue: 0, green: 3},
				{red: 14, blue: 15, green: 3},
			},
		},
		{
			id: 5,
			sets: []Set{
				{red: 6, blue: 1, green: 3},
				{red: 1, blue: 2, green: 2},
			},
		},
	}
	actual := parsePuzzleInput(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("solvePartOne(%v) = %d; want %d", input, actual, expected)
	}
}

func TestSolvePartOne(t *testing.T) {
	input := []Game{
		{
			id: 1,
			sets: []Set{
				{red: 4, blue: 3, green: 0},
				{red: 1, blue: 6, green: 2},
				{red: 0, blue: 0, green: 2},
			},
		},
		{
			id: 2,
			sets: []Set{
				{red: 0, blue: 1, green: 2},
				{red: 1, blue: 4, green: 3},
				{red: 0, blue: 1, green: 1},
			},
		},
		{
			id: 3,
			sets: []Set{
				{red: 20, blue: 6, green: 8},
				{red: 4, blue: 5, green: 13},
				{red: 1, blue: 0, green: 5},
			},
		},
		{
			id: 4,
			sets: []Set{
				{red: 3, blue: 6, green: 1},
				{red: 6, blue: 0, green: 3},
				{red: 14, blue: 15, green: 3},
			},
		},
		{
			id: 5,
			sets: []Set{
				{red: 6, blue: 1, green: 3},
				{red: 1, blue: 2, green: 2},
			},
		},
	}
	expected := 8
	actual := solvePartOne(input)

	if actual != expected {
		t.Errorf("solvePartOne(%v) = %d; want %d", input, actual, expected)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input := []Game{
		{
			id: 1,
			sets: []Set{
				{red: 4, blue: 3, green: 0},
				{red: 1, blue: 6, green: 2},
				{red: 0, blue: 0, green: 2},
			},
		},
		{
			id: 2,
			sets: []Set{
				{red: 0, blue: 1, green: 2},
				{red: 1, blue: 4, green: 3},
				{red: 0, blue: 1, green: 1},
			},
		},
		{
			id: 3,
			sets: []Set{
				{red: 20, blue: 6, green: 8},
				{red: 4, blue: 5, green: 13},
				{red: 1, blue: 0, green: 5},
			},
		},
		{
			id: 4,
			sets: []Set{
				{red: 3, blue: 6, green: 1},
				{red: 6, blue: 0, green: 3},
				{red: 14, blue: 15, green: 3},
			},
		},
		{
			id: 5,
			sets: []Set{
				{red: 6, blue: 1, green: 3},
				{red: 1, blue: 2, green: 2},
			},
		},
	}
	expected := 2286
	actual := solvePartTwo(input)

	if actual != expected {
		t.Errorf("solvePartTwo(%v) = %d; want %d", input, actual, expected)
	}
}
