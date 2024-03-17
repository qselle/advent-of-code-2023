package main

import (
	"fmt"
	"strings"

	"github.com/quentinselle/aoc/2023/go/utils"
)

type Race struct {
	Time           int
	DistanceRecord int
}

type Races []Race

func (r Race) WayToBreakRecords() int {
	win := 0
	for elapsed := 0; elapsed < r.Time; elapsed++ {
		if elapsed*(r.Time-elapsed) > r.DistanceRecord {
			win++
		}
	}
	return win
}

func (r Races) ComputeTotalWayToBreakRecords() int {
	total := 1
	for _, race := range r {
		total *= race.WayToBreakRecords()
	}
	return total
}

func parseRace(input []string) Races {
	races := Races{}
	for _, time := range utils.ExtractNumbersFromString(strings.TrimPrefix(input[0], "Time:")) {
		races = append(races, Race{Time: time})
	}
	for index, distance := range utils.ExtractNumbersFromString(strings.TrimPrefix(input[1], "Distance:")) {
		races[index].DistanceRecord = distance
	}
	return races
}

func parseRaceWithNoKerning(input []string) Race {
	return Race{
		utils.ExtractNumberFromStringIgnoringSpaces(strings.TrimPrefix(input[0], "Time:")),
		utils.ExtractNumberFromStringIgnoringSpaces(strings.TrimPrefix(input[1], "Distance:")),
	}
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	races := parseRace(input)
	fmt.Println("Part 1", races.ComputeTotalWayToBreakRecords())
	race := parseRaceWithNoKerning(input)
	fmt.Println("Part 2", race.WayToBreakRecords())
}
