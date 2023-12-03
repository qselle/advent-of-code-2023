package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readFileByLine(inputFile string) []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

type Position struct {
	x int
	y int
}

type SchemeSymbols map[Position]rune

func (s SchemeSymbols) isPositionSymbol(currentPosition Position) bool {
	_, found := s[currentPosition]
	return found
}

func (s SchemeSymbols) isPositionAdjacent(currentPosition Position) bool {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if s.isPositionSymbol(Position{
				currentPosition.x + x,
				currentPosition.y + y,
			}) {
				return true
			}
		}
	}
	return false
}

func (s SchemeSymbols) isPositionSymbolGear(currentPosition Position) bool {
	symbol, found := s[currentPosition]
	return found && symbol == '*'
}

func (s SchemeSymbols) isPositionAdjacentToGear(currentPosition Position) (bool, Position) {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if s.isPositionSymbolGear(Position{
				currentPosition.x + x,
				currentPosition.y + y,
			}) {
				return true, Position{
					currentPosition.x + x,
					currentPosition.y + y,
				}
			}
		}
	}
	return false, Position{}
}

func parseSymbols(input []string) SchemeSymbols {
	engineSchemeSymbols := make(SchemeSymbols)
	for y, line := range input {
		for x, sym := range line {
			if sym != '.' && !unicode.IsNumber(sym) {
				engineSchemeSymbols[Position{x: x, y: y}] = sym
			}
		}
	}
	return engineSchemeSymbols
}

func solvePartOne(input []string) int {
	sum := 0
	engineSchemeSymbol := parseSymbols(input)
	for y, line := range input {
		adjacent := false
		number := 0
		for x, sym := range line {
			if unicode.IsNumber(sym) {
				if engineSchemeSymbol.isPositionAdjacent(Position{x, y}) {
					adjacent = true
				}
				number = number*10 + int(sym-'0')
				continue
			}
			if number > 0 && adjacent {
				sum += number
			}
			adjacent = false
			number = 0
		}
		if number > 0 && adjacent {
			sum += number
		}
	}
	return sum
}

type PossibleGear struct {
	ref   int
	ratio int
}

type SchemePossibleGears map[Position]PossibleGear

func solvePartTwo(input []string) int {
	sum := 0
	engineSchemeSymbol := parseSymbols(input)
	possibleGear := make(SchemePossibleGears)
	for y, line := range input {
		adjacent := false
		number := 0
		gearPosition := Position{}
		for x, sym := range line {
			if unicode.IsNumber(sym) {
				found, possibleGearPosition := engineSchemeSymbol.isPositionAdjacentToGear(Position{x, y})
				if found {
					gearPosition = possibleGearPosition
					adjacent = true
				}
				number = number*10 + int(sym-'0')
				continue
			}
			if number > 0 && adjacent {
				gear, found := possibleGear[gearPosition]
				if !found {
					possibleGear[gearPosition] = PossibleGear{
						ref:   1,
						ratio: number,
					}
				} else {
					gear.ref++
					gear.ratio *= number
					possibleGear[gearPosition] = gear
				}
			}
			adjacent = false
			number = 0
		}
		if number > 0 && adjacent {
			gear, found := possibleGear[gearPosition]
			if !found {
				possibleGear[gearPosition] = PossibleGear{
					ref:   1,
					ratio: number,
				}
			} else {
				gear.ref++
				gear.ratio *= number
				possibleGear[gearPosition] = gear
			}
		}
	}
	for _, v := range possibleGear {
		if v.ref == 2 {
			sum += v.ratio
		}
	}
	return sum
}

func main() {
	input := readFileByLine("input.txt")
	fmt.Println("Part 1:", solvePartOne(input))
	fmt.Println("Part 2:", solvePartTwo(input))
}
