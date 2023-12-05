package main

import (
	"reflect"
	"testing"
)

func TestParsePuzzleInput(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	expected := ScratchCards{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}

	actual := parseScratchCards(input)

	for i, card := range actual {
		if !reflect.DeepEqual(card, expected[i]) {
			t.Errorf("Expected %v, got %v", expected[i], card)
		}
	}
}

func TestTotalPointsWorthByCard(t *testing.T) {
	scratchCardTests := []struct {
		Expected int
		Card     ScratchCard
	}{
		{
			Expected: 8,
			Card: ScratchCard{
				WinningNumbers: []int{41, 48, 83, 86, 17},
				CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
		{
			Expected: 2,
			Card: ScratchCard{
				WinningNumbers: []int{13, 32, 20, 16, 61},
				CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
		},
		{
			Expected: 2,
			Card: ScratchCard{
				WinningNumbers: []int{1, 21, 53, 59, 44},
				CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
		},
		{
			Expected: 1,
			Card: ScratchCard{
				WinningNumbers: []int{41, 92, 73, 84, 69},
				CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
		},
		{
			Expected: 0,
			Card: ScratchCard{
				WinningNumbers: []int{87, 83, 26, 28, 32},
				CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
		},
		{
			Expected: 0,
			Card: ScratchCard{
				WinningNumbers: []int{31, 18, 13, 56, 72},
				CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
		},
	}
	for _, test := range scratchCardTests {
		actual := test.Card.TotalPointsWorth()
		if actual != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, actual)
		}
	}
}

func TestTotalMatches(t *testing.T) {
	scratchCardTests := []struct {
		Expected int
		Card     ScratchCard
	}{
		{
			Expected: 4,
			Card: ScratchCard{
				WinningNumbers: []int{41, 48, 83, 86, 17},
				CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
		{
			Expected: 2,
			Card: ScratchCard{
				WinningNumbers: []int{13, 32, 20, 16, 61},
				CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
		},
		{
			Expected: 2,
			Card: ScratchCard{
				WinningNumbers: []int{1, 21, 53, 59, 44},
				CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
		},
		{
			Expected: 1,
			Card: ScratchCard{
				WinningNumbers: []int{41, 92, 73, 84, 69},
				CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
		},
		{
			Expected: 0,
			Card: ScratchCard{
				WinningNumbers: []int{87, 83, 26, 28, 32},
				CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
		},
		{
			Expected: 0,
			Card: ScratchCard{
				WinningNumbers: []int{31, 18, 13, 56, 72},
				CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
		},
	}
	for _, test := range scratchCardTests {
		actual := test.Card.TotalMatches()
		if actual != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, actual)
		}
	}
}

func TestTotalPointsWorth(t *testing.T) {
	scratchCards := ScratchCards{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}
	expected := 13
	actual := scratchCards.TotalPointsWorth()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTotalScratchCardsWin(t *testing.T) {
	scratchCards := ScratchCards{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}
	expected := 30
	actual := scratchCards.TotalScratchCardsWin(0, len(scratchCards), len(scratchCards))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkScratchCard_TotalPointsWorth(b *testing.B) {
	scratchCards := ScratchCards{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}
	for i := 0; i < b.N; i++ {
		scratchCards.TotalPointsWorth()
	}
}

func BenchmarkScratchCard_TotalScratchCardsWin(b *testing.B) {
	scratchCards := ScratchCards{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			CardNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			CardNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			CardNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			CardNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			CardNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}
	for i := 0; i < b.N; i++ {
		scratchCards.TotalScratchCardsWin(0, len(scratchCards), len(scratchCards))
	}
}
