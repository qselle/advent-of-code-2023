package main

import (
	"fmt"
	"github.com/quentinselle/aoc/2023/utils"
	"strings"
)

type Rule struct {
	DestinationRangeStart int
	SourceRangeStart      int
	Length                int
}

type Map struct {
	From  string
	To    string
	Rules []Rule
}

type Maps []Map

type Almanac struct {
	Seeds []int
	Maps  Maps
}

func (m Maps) SeedToLocation(seed int) int {
nextMap:
	for _, currMap := range m {
		for _, currRule := range currMap.Rules {
			start := currRule.SourceRangeStart
			end := currRule.SourceRangeStart + currRule.Length
			dest := currRule.DestinationRangeStart
			if seed >= start && seed < end {
				seed = dest + (seed - start)
				continue nextMap
			}
		}
	}
	return seed
}

func (al Almanac) LowestLocationNumberFromSeed() int {
	actual := 0
	lowest := 0
	for _, seed := range al.Seeds {
		actual = al.Maps.SeedToLocation(seed)
		if lowest == 0 {
			lowest = actual
		}
		if actual < lowest {
			lowest = actual
		}
	}
	return lowest
}

func (al Almanac) LowestLocationNumberFromSeedRange() int {
	actual := 0
	lowest := 0

	for index := 0; index < len(al.Seeds); index += 2 {
		for seed := al.Seeds[index]; seed < al.Seeds[index]+al.Seeds[index+1]; seed++ {
			actual = al.Maps.SeedToLocation(seed)
			if lowest == 0 {
				lowest = actual
			}
			if actual < lowest {
				lowest = actual
			}
		}
	}
	return lowest
}

func parseAlmanac(input []string) Almanac {
	almanac := Almanac{
		Seeds: utils.ExtractNumbersFromString(strings.Split(input[0], ":")[1]),
	}
	var processingMap Map

	for _, line := range input[2:] {
		if strings.HasSuffix(line, "map:") {
			mapName := strings.Split(strings.Fields(line)[0], "-to-")
			processingMap.From, processingMap.To = mapName[0], mapName[1]
			continue
		}
		if line == "" {
			almanac.Maps = append(almanac.Maps, processingMap)
			processingMap = Map{}
			continue
		}
		numbers := utils.ExtractNumbersFromString(line)
		rule := Rule{numbers[0], numbers[1], numbers[2]}
		processingMap.Rules = append(processingMap.Rules, rule)
	}
	// last elem
	almanac.Maps = append(almanac.Maps, processingMap)

	return almanac
}

func main() {
	al := parseAlmanac(utils.ReadFileByLine("input.txt"))
	fmt.Println("Part 1:", al.LowestLocationNumberFromSeed())
	fmt.Println("Part 2:", al.LowestLocationNumberFromSeedRange())
}
