package main

import (
	"testing"
)

func TestExtrapolateBackwardsReport(t *testing.T) {
	test := []struct {
		Expected int
		Input    []string
	}{
		{
			Expected: 4,
			Input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
		},
		{
			Expected: 8,
			Input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
		},
	}
	for _, test := range test {
		tiles := parseTiles(test.Input)
		actual := tiles.ComputeFarthestStep()
		if actual != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, actual)
		}
	}
}
