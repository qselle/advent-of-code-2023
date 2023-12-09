package main

import (
	"reflect"
	"testing"
)

func TestParseMaps(t *testing.T) {
	input := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	expected := Maps{
		Directions: []string{"R", "L"},
		Nodes: map[string]Node{
			"AAA": {
				Left:  "BBB",
				Right: "CCC",
			},
			"BBB": {
				Left:  "DDD",
				Right: "EEE",
			},
			"CCC": {
				Left:  "ZZZ",
				Right: "GGG",
			},
			"DDD": {
				Left:  "DDD",
				Right: "DDD",
			},
			"EEE": {
				Left:  "EEE",
				Right: "EEE",
			},
			"GGG": {
				Left:  "GGG",
				Right: "GGG",
			},
			"ZZZ": {
				Left:  "ZZZ",
				Right: "ZZZ",
			},
		},
	}
	actual := parseMaps(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestStepsToReachEnd(t *testing.T) {

	input := []struct {
		Expected int
		Maps     Maps
	}{
		{
			Expected: 2,
			Maps: Maps{
				Directions: []string{"R", "L"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "CCC",
					},
					"BBB": {
						Left:  "DDD",
						Right: "EEE",
					},
					"CCC": {
						Left:  "ZZZ",
						Right: "GGG",
					},
					"DDD": {
						Left:  "DDD",
						Right: "DDD",
					},
					"EEE": {
						Left:  "EEE",
						Right: "EEE",
					},
					"GGG": {
						Left:  "GGG",
						Right: "GGG",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
		{
			Expected: 6,
			Maps: Maps{
				Directions: []string{"L", "L", "R"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "BBB",
					},
					"BBB": {
						Left:  "AAA",
						Right: "ZZZ",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
	}
	for _, test := range input {
		actual := test.Maps.StepsToReachEnd()
		if actual != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, actual)
		}
	}
}
