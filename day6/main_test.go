package main

import (
	"reflect"
	"testing"
)

func TestParseRace(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expected := Races{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	actual := parseRace(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseRaceWithNoKerning(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expected := Race{
		71530,
		940200,
	}
	actual := parseRaceWithNoKerning(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestRace_WayToBreakRecordsByRace(t *testing.T) {
	input := []struct {
		Expected int
		Race     Race
	}{
		{
			Expected: 4,
			Race:     Race{7, 9},
		},
		{
			Expected: 8,
			Race:     Race{15, 40},
		},
		{
			Expected: 9,
			Race:     Race{30, 200},
		},
	}

	for _, r := range input {
		actual := r.Race.WayToBreakRecords()
		expected := r.Expected
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestRaces_ComputeTotalWayToBreakRecords(t *testing.T) {
	input := Races{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	expected := 288
	actual := input.ComputeTotalWayToBreakRecords()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkRaces_ComputeTotalWayToBreakRecords(b *testing.B) {
	input := Races{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	for i := 0; i < b.N; i++ {
		input.ComputeTotalWayToBreakRecords()
	}
}
func BenchmarkRace_WayToBreakRecordsWithNoKerning(b *testing.B) {
	input := Race{
		71530,
		940200,
	}
	for i := 0; i < b.N; i++ {
		input.WayToBreakRecords()
	}
}
