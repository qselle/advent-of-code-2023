package main

import (
	"reflect"
	"testing"
)

func TestParsing(t *testing.T) {
	input := []string{"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	expected := Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: []Map{
			{
				From: "seed",
				To:   "soil",
				Rules: []Rule{
					{
						DestinationRangeStart: 50,
						SourceRangeStart:      98,
						Length:                2,
					},
					{
						DestinationRangeStart: 52,
						SourceRangeStart:      50,
						Length:                48,
					},
				},
			},
			{
				From: "soil",
				To:   "fertilizer",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      15,
						Length:                37,
					},
					{
						DestinationRangeStart: 37,
						SourceRangeStart:      52,
						Length:                2,
					},
					{
						DestinationRangeStart: 39,
						SourceRangeStart:      0,
						Length:                15,
					},
				},
			},
			{
				From: "fertilizer",
				To:   "water",
				Rules: []Rule{
					{
						DestinationRangeStart: 49,
						SourceRangeStart:      53,
						Length:                8,
					},
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      11,
						Length:                42,
					},
					{
						DestinationRangeStart: 42,
						SourceRangeStart:      0,
						Length:                7,
					},
					{
						DestinationRangeStart: 57,
						SourceRangeStart:      7,
						Length:                4,
					},
				},
			},
			{
				From: "water",
				To:   "light",
				Rules: []Rule{
					{
						DestinationRangeStart: 88,
						SourceRangeStart:      18,
						Length:                7,
					},
					{
						DestinationRangeStart: 18,
						SourceRangeStart:      25,
						Length:                70,
					},
				},
			},
			{
				From: "light",
				To:   "temperature",
				Rules: []Rule{
					{
						DestinationRangeStart: 45,
						SourceRangeStart:      77,
						Length:                23,
					},
					{
						DestinationRangeStart: 81,
						SourceRangeStart:      45,
						Length:                19,
					},
					{
						DestinationRangeStart: 68,
						SourceRangeStart:      64,
						Length:                13,
					},
				},
			},
			{
				From: "temperature",
				To:   "humidity",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      69,
						Length:                1,
					},
					{
						DestinationRangeStart: 1,
						SourceRangeStart:      0,
						Length:                69,
					},
				},
			},
			{
				From: "humidity",
				To:   "location",
				Rules: []Rule{
					{
						DestinationRangeStart: 60,
						SourceRangeStart:      56,
						Length:                37,
					},
					{
						DestinationRangeStart: 56,
						SourceRangeStart:      93,
						Length:                4,
					},
				},
			},
		},
	}
	actual := parseAlmanac(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAlmanac_LowestLocationNumberFromSeed(t *testing.T) {
	al := Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: []Map{
			{
				From: "seed",
				To:   "soil",
				Rules: []Rule{
					{
						DestinationRangeStart: 50,
						SourceRangeStart:      98,
						Length:                2,
					},
					{
						DestinationRangeStart: 52,
						SourceRangeStart:      50,
						Length:                48,
					},
				},
			},
			{
				From: "soil",
				To:   "fertilizer",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      15,
						Length:                37,
					},
					{
						DestinationRangeStart: 37,
						SourceRangeStart:      52,
						Length:                2,
					},
					{
						DestinationRangeStart: 39,
						SourceRangeStart:      0,
						Length:                15,
					},
				},
			},
			{
				From: "fertilizer",
				To:   "water",
				Rules: []Rule{
					{
						DestinationRangeStart: 49,
						SourceRangeStart:      53,
						Length:                8,
					},
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      11,
						Length:                42,
					},
					{
						DestinationRangeStart: 42,
						SourceRangeStart:      0,
						Length:                7,
					},
					{
						DestinationRangeStart: 57,
						SourceRangeStart:      7,
						Length:                4,
					},
				},
			},
			{
				From: "water",
				To:   "light",
				Rules: []Rule{
					{
						DestinationRangeStart: 88,
						SourceRangeStart:      18,
						Length:                7,
					},
					{
						DestinationRangeStart: 18,
						SourceRangeStart:      25,
						Length:                70,
					},
				},
			},
			{
				From: "light",
				To:   "temperature",
				Rules: []Rule{
					{
						DestinationRangeStart: 45,
						SourceRangeStart:      77,
						Length:                23,
					},
					{
						DestinationRangeStart: 81,
						SourceRangeStart:      45,
						Length:                19,
					},
					{
						DestinationRangeStart: 68,
						SourceRangeStart:      64,
						Length:                13,
					},
				},
			},
			{
				From: "temperature",
				To:   "humidity",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      69,
						Length:                1,
					},
					{
						DestinationRangeStart: 1,
						SourceRangeStart:      0,
						Length:                69,
					},
				},
			},
			{
				From: "humidity",
				To:   "location",
				Rules: []Rule{
					{
						DestinationRangeStart: 60,
						SourceRangeStart:      56,
						Length:                37,
					},
					{
						DestinationRangeStart: 56,
						SourceRangeStart:      93,
						Length:                4,
					},
				},
			},
		},
	}
	expected := 35
	actual := al.LowestLocationNumberFromSeed()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestAlmanac_LowestLocationNumberFromSeedRange(t *testing.T) {
	al := Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: []Map{
			{
				From: "seed",
				To:   "soil",
				Rules: []Rule{
					{
						DestinationRangeStart: 50,
						SourceRangeStart:      98,
						Length:                2,
					},
					{
						DestinationRangeStart: 52,
						SourceRangeStart:      50,
						Length:                48,
					},
				},
			},
			{
				From: "soil",
				To:   "fertilizer",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      15,
						Length:                37,
					},
					{
						DestinationRangeStart: 37,
						SourceRangeStart:      52,
						Length:                2,
					},
					{
						DestinationRangeStart: 39,
						SourceRangeStart:      0,
						Length:                15,
					},
				},
			},
			{
				From: "fertilizer",
				To:   "water",
				Rules: []Rule{
					{
						DestinationRangeStart: 49,
						SourceRangeStart:      53,
						Length:                8,
					},
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      11,
						Length:                42,
					},
					{
						DestinationRangeStart: 42,
						SourceRangeStart:      0,
						Length:                7,
					},
					{
						DestinationRangeStart: 57,
						SourceRangeStart:      7,
						Length:                4,
					},
				},
			},
			{
				From: "water",
				To:   "light",
				Rules: []Rule{
					{
						DestinationRangeStart: 88,
						SourceRangeStart:      18,
						Length:                7,
					},
					{
						DestinationRangeStart: 18,
						SourceRangeStart:      25,
						Length:                70,
					},
				},
			},
			{
				From: "light",
				To:   "temperature",
				Rules: []Rule{
					{
						DestinationRangeStart: 45,
						SourceRangeStart:      77,
						Length:                23,
					},
					{
						DestinationRangeStart: 81,
						SourceRangeStart:      45,
						Length:                19,
					},
					{
						DestinationRangeStart: 68,
						SourceRangeStart:      64,
						Length:                13,
					},
				},
			},
			{
				From: "temperature",
				To:   "humidity",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      69,
						Length:                1,
					},
					{
						DestinationRangeStart: 1,
						SourceRangeStart:      0,
						Length:                69,
					},
				},
			},
			{
				From: "humidity",
				To:   "location",
				Rules: []Rule{
					{
						DestinationRangeStart: 60,
						SourceRangeStart:      56,
						Length:                37,
					},
					{
						DestinationRangeStart: 56,
						SourceRangeStart:      93,
						Length:                4,
					},
				},
			},
		},
	}
	expected := 46
	actual := al.LowestLocationNumberFromSeedRange()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkAlmanac_LowestLocationNumberFromSeed(b *testing.B) {
	al := Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: []Map{
			{
				From: "seed",
				To:   "soil",
				Rules: []Rule{
					{
						DestinationRangeStart: 50,
						SourceRangeStart:      98,
						Length:                2,
					},
					{
						DestinationRangeStart: 52,
						SourceRangeStart:      50,
						Length:                48,
					},
				},
			},
			{
				From: "soil",
				To:   "fertilizer",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      15,
						Length:                37,
					},
					{
						DestinationRangeStart: 37,
						SourceRangeStart:      52,
						Length:                2,
					},
					{
						DestinationRangeStart: 39,
						SourceRangeStart:      0,
						Length:                15,
					},
				},
			},
			{
				From: "fertilizer",
				To:   "water",
				Rules: []Rule{
					{
						DestinationRangeStart: 49,
						SourceRangeStart:      53,
						Length:                8,
					},
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      11,
						Length:                42,
					},
					{
						DestinationRangeStart: 42,
						SourceRangeStart:      0,
						Length:                7,
					},
					{
						DestinationRangeStart: 57,
						SourceRangeStart:      7,
						Length:                4,
					},
				},
			},
			{
				From: "water",
				To:   "light",
				Rules: []Rule{
					{
						DestinationRangeStart: 88,
						SourceRangeStart:      18,
						Length:                7,
					},
					{
						DestinationRangeStart: 18,
						SourceRangeStart:      25,
						Length:                70,
					},
				},
			},
			{
				From: "light",
				To:   "temperature",
				Rules: []Rule{
					{
						DestinationRangeStart: 45,
						SourceRangeStart:      77,
						Length:                23,
					},
					{
						DestinationRangeStart: 81,
						SourceRangeStart:      45,
						Length:                19,
					},
					{
						DestinationRangeStart: 68,
						SourceRangeStart:      64,
						Length:                13,
					},
				},
			},
			{
				From: "temperature",
				To:   "humidity",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      69,
						Length:                1,
					},
					{
						DestinationRangeStart: 1,
						SourceRangeStart:      0,
						Length:                69,
					},
				},
			},
			{
				From: "humidity",
				To:   "location",
				Rules: []Rule{
					{
						DestinationRangeStart: 60,
						SourceRangeStart:      56,
						Length:                37,
					},
					{
						DestinationRangeStart: 56,
						SourceRangeStart:      93,
						Length:                4,
					},
				},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		al.LowestLocationNumberFromSeed()
	}
}

func BenchmarkAlmanac_LowestLocationNumberFromSeedRange(b *testing.B) {
	al := Almanac{
		Seeds: []int{79, 14, 55, 13},
		Maps: []Map{
			{
				From: "seed",
				To:   "soil",
				Rules: []Rule{
					{
						DestinationRangeStart: 50,
						SourceRangeStart:      98,
						Length:                2,
					},
					{
						DestinationRangeStart: 52,
						SourceRangeStart:      50,
						Length:                48,
					},
				},
			},
			{
				From: "soil",
				To:   "fertilizer",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      15,
						Length:                37,
					},
					{
						DestinationRangeStart: 37,
						SourceRangeStart:      52,
						Length:                2,
					},
					{
						DestinationRangeStart: 39,
						SourceRangeStart:      0,
						Length:                15,
					},
				},
			},
			{
				From: "fertilizer",
				To:   "water",
				Rules: []Rule{
					{
						DestinationRangeStart: 49,
						SourceRangeStart:      53,
						Length:                8,
					},
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      11,
						Length:                42,
					},
					{
						DestinationRangeStart: 42,
						SourceRangeStart:      0,
						Length:                7,
					},
					{
						DestinationRangeStart: 57,
						SourceRangeStart:      7,
						Length:                4,
					},
				},
			},
			{
				From: "water",
				To:   "light",
				Rules: []Rule{
					{
						DestinationRangeStart: 88,
						SourceRangeStart:      18,
						Length:                7,
					},
					{
						DestinationRangeStart: 18,
						SourceRangeStart:      25,
						Length:                70,
					},
				},
			},
			{
				From: "light",
				To:   "temperature",
				Rules: []Rule{
					{
						DestinationRangeStart: 45,
						SourceRangeStart:      77,
						Length:                23,
					},
					{
						DestinationRangeStart: 81,
						SourceRangeStart:      45,
						Length:                19,
					},
					{
						DestinationRangeStart: 68,
						SourceRangeStart:      64,
						Length:                13,
					},
				},
			},
			{
				From: "temperature",
				To:   "humidity",
				Rules: []Rule{
					{
						DestinationRangeStart: 0,
						SourceRangeStart:      69,
						Length:                1,
					},
					{
						DestinationRangeStart: 1,
						SourceRangeStart:      0,
						Length:                69,
					},
				},
			},
			{
				From: "humidity",
				To:   "location",
				Rules: []Rule{
					{
						DestinationRangeStart: 60,
						SourceRangeStart:      56,
						Length:                37,
					},
					{
						DestinationRangeStart: 56,
						SourceRangeStart:      93,
						Length:                4,
					},
				},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		al.LowestLocationNumberFromSeedRange()
	}
}
