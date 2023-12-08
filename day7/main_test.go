package main

import (
	"reflect"
	"testing"
)

func TestParseHands(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expected := Hands{
		{
			Cards: []Card{
				{Value: "3", Strength: 2},
				{Value: "2", Strength: 1},
				{Value: "T", Strength: 9},
				{Value: "3", Strength: 2},
				{Value: "K", Strength: 12},
			},
			Bid:  765,
			Type: OnePair,
		},
		{
			Cards: []Card{
				{Value: "T", Strength: 9},
				{Value: "5", Strength: 4},
				{Value: "5", Strength: 4},
				{Value: "J", Strength: 10},
				{Value: "5", Strength: 4},
			},
			Bid:  684,
			Type: ThreeOfAKind,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "K", Strength: 12},
				{Value: "6", Strength: 5},
				{Value: "7", Strength: 6},
				{Value: "7", Strength: 6},
			},
			Bid:  28,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "T", Strength: 9},
				{Value: "J", Strength: 10},
				{Value: "J", Strength: 10},
				{Value: "T", Strength: 9},
			},
			Bid:  220,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "J", Strength: 10},
				{Value: "A", Strength: 13},
			},
			Bid:  483,
			Type: ThreeOfAKind,
		},
	}
	actual := parseHands(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSortHands(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expected := Hands{
		{
			Cards: []Card{
				{Value: "3", Strength: 2},
				{Value: "2", Strength: 1},
				{Value: "T", Strength: 9},
				{Value: "3", Strength: 2},
				{Value: "K", Strength: 12},
			},
			Bid:  765,
			Type: OnePair,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "T", Strength: 9},
				{Value: "J", Strength: 10},
				{Value: "J", Strength: 10},
				{Value: "T", Strength: 9},
			},
			Bid:  220,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "K", Strength: 12},
				{Value: "6", Strength: 5},
				{Value: "7", Strength: 6},
				{Value: "7", Strength: 6},
			},
			Bid:  28,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "T", Strength: 9},
				{Value: "5", Strength: 4},
				{Value: "5", Strength: 4},
				{Value: "J", Strength: 10},
				{Value: "5", Strength: 4},
			},
			Bid:  684,
			Type: ThreeOfAKind,
		},
		{
			Cards: []Card{
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "J", Strength: 10},
				{Value: "A", Strength: 13},
			},
			Bid:  483,
			Type: ThreeOfAKind,
		},
	}
	actual := parseHands(input)
	actual.SortDesc()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestComputeTotalWinnings(t *testing.T) {
	input := Hands{
		{
			Cards: []Card{
				{Value: "3", Strength: 2},
				{Value: "2", Strength: 1},
				{Value: "T", Strength: 9},
				{Value: "3", Strength: 2},
				{Value: "K", Strength: 12},
			},
			Bid:  765,
			Type: OnePair,
		},
		{
			Cards: []Card{
				{Value: "T", Strength: 9},
				{Value: "5", Strength: 4},
				{Value: "5", Strength: 4},
				{Value: "J", Strength: 10},
				{Value: "5", Strength: 4},
			},
			Bid:  684,
			Type: ThreeOfAKind,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "K", Strength: 12},
				{Value: "6", Strength: 5},
				{Value: "7", Strength: 6},
				{Value: "7", Strength: 6},
			},
			Bid:  28,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "K", Strength: 12},
				{Value: "T", Strength: 9},
				{Value: "J", Strength: 10},
				{Value: "J", Strength: 10},
				{Value: "T", Strength: 9},
			},
			Bid:  220,
			Type: TwoPair,
		},
		{
			Cards: []Card{
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "Q", Strength: 11},
				{Value: "J", Strength: 10},
				{Value: "A", Strength: 13},
			},
			Bid:  483,
			Type: ThreeOfAKind,
		},
	}
	expected := 6440
	input.SortDesc()
	actual := input.ComputeTotalWinnings()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
