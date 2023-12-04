package main

import (
	"fmt"
	"github.com/quentinselle/aoc/2023/utils"
	"unicode"
)

type Coordinates struct {
	Y int
	X int
}

type symbol string
type Symbols map[Coordinates]symbol

type Part struct {
	Number      int
	Coordinates []Coordinates
}

func (p Part) IsCoordinatesAdjacent(symbols Symbols) (bool, Coordinates) {
	for _, coordinates := range p.Coordinates {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if _, found := symbols[Coordinates{coordinates.Y + y, coordinates.X + x}]; found {
					return true, Coordinates{coordinates.Y + y, coordinates.X + x}
				}
			}
		}
	}
	return false, Coordinates{}
}

type PossibleGear struct {
	refs  int
	ratio int
}

type PossibleGears map[Coordinates]PossibleGear

type EngineSchematic struct {
	Parts   []Part
	Symbols Symbols
}

func (es EngineSchematic) PartNumbersSum() int {
	sum := 0
	for _, part := range es.Parts {
		if found, _ := part.IsCoordinatesAdjacent(es.Symbols); found {
			sum += part.Number
		}
	}
	return sum
}

func (es EngineSchematic) GearRatiosSum() int {
	sum := 0

	// keep "*" symbols only, so IsCoordinatesAdjacent will only match those.
	for k, v := range es.Symbols {
		if v != "*" {
			delete(es.Symbols, k)
		}
	}

	possibleGears := make(PossibleGears)
	for _, part := range es.Parts {
		if found, coordinates := part.IsCoordinatesAdjacent(es.Symbols); found {
			possibleGear, found := possibleGears[coordinates]
			if !found {
				possibleGear = PossibleGear{1, part.Number}
			} else {
				possibleGear.refs++
				possibleGear.ratio *= part.Number
			}
			possibleGears[coordinates] = possibleGear
		}
	}

	for _, v := range possibleGears {
		if v.refs == 2 {
			sum += v.ratio
		}
	}
	return sum
}

func ParseEngineSchematic(input []string) *EngineSchematic {
	es := &EngineSchematic{}
	es.Symbols = make(Symbols)

	for y, line := range input {
		for x, char := range line {
			if char != '.' && !unicode.IsNumber(char) {
				es.Symbols[Coordinates{y, x}] = symbol(char)
			}
		}
	}

	for y, line := range input {
		part := Part{}
		for x, char := range line {
			if unicode.IsNumber(char) {
				part.Number = part.Number*10 + utils.RuneToInt(char)
				part.Coordinates = append(part.Coordinates, Coordinates{y, x})
				continue
			}
			if part.Number > 0 {
				es.Parts = append(es.Parts, part)
				part = Part{}
			}
		}
		if part.Number > 0 {
			es.Parts = append(es.Parts, part)
			part = Part{}
		}
	}
	return es
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	es := ParseEngineSchematic(input)
	fmt.Println("Part 1:", es.PartNumbersSum())
	fmt.Println("Part 2:", es.GearRatiosSum())
}
