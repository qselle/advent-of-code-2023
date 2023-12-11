package main

import (
	"reflect"
	"testing"
)

func TestParseReport(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	expected := Report{
		Histories: []History{
			{0, 3, 6, 9, 12, 15},
			{1, 3, 6, 10, 15, 21},
			{10, 13, 16, 21, 30, 45},
		},
	}
	actual := parseHistories(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestExtrapolateReport(t *testing.T) {
	input := Report{
		Histories: []History{
			{0, 3, 6, 9, 12, 15},
			{1, 3, 6, 10, 15, 21},
			{10, 13, 16, 21, 30, 45},
		},
	}
	expected := 114
	actual := input.Extrapolate()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkExtrapolateReport(b *testing.B) {
	input := Report{
		Histories: []History{
			{0, 3, 6, 9, 12, 15},
			{1, 3, 6, 10, 15, 21},
			{10, 13, 16, 21, 30, 45},
		},
	}
	for i := 0; i < b.N; i++ {
		input.Extrapolate()
	}
}

func TestExtrapolateBackwardsReport(t *testing.T) {
	input := Report{
		Histories: []History{
			{0, 3, 6, 9, 12, 15},
			{1, 3, 6, 10, 15, 21},
			{10, 13, 16, 21, 30, 45},
		},
	}
	expected := 2
	actual := input.ExtrapolateBackwards()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkExtrapolateBackwardsReport(b *testing.B) {
	input := Report{
		Histories: []History{
			{0, 3, 6, 9, 12, 15},
			{1, 3, 6, 10, 15, 21},
			{10, 13, 16, 21, 30, 45},
		},
	}
	for i := 0; i < b.N; i++ {
		input.ExtrapolateBackwards()
	}
}
