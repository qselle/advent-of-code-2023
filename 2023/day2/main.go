package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsePuzzleInput(inputFile string) []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func solvePartTwo(input []string) (total int) {
	for gameIndex, line := range input {
		line = strings.TrimPrefix(line, fmt.Sprintf("Game %d:", gameIndex+1))
		redCubesMin, greenCubesMin, blueCubesMin := 0, 0, 0
		for _, set := range strings.Split(line, ";") {
			redCubes, greenCubes, blueCubes := 0, 0, 0
			for _, subset := range strings.Split(set, ",") {
				currentCube := strings.Fields(subset)
				cubesNumber, err := strconv.Atoi(currentCube[0])
				if err != nil {
					panic(err)
				}
				switch currentCube[1] {
				case "red":
					redCubes += cubesNumber
				case "green":
					greenCubes += cubesNumber
				case "blue":
					blueCubes += cubesNumber
				default:
					panic("Invalid cube color")
				}
			}
			if redCubes > redCubesMin {
				redCubesMin = redCubes
			}
			if greenCubes > greenCubesMin {
				greenCubesMin = greenCubes
			}
			if blueCubes > blueCubesMin {
				blueCubesMin = blueCubes
			}
		}
		total += redCubesMin * greenCubesMin * blueCubesMin
	}
	return
}

func solvePartOne(input []string) (total int) {
	const (
		redCubesInBag   = 12
		greenCubesInBag = 13
		blueCubesInBag  = 14
	)

nextGame:
	for gameIndex, line := range input {
		line = strings.TrimPrefix(line, fmt.Sprintf("Game %d:", gameIndex+1))

		for _, set := range strings.Split(line, ";") {
			redCubes, greenCubes, blueCubes := redCubesInBag, greenCubesInBag, blueCubesInBag
			for _, subset := range strings.Split(set, ",") {
				currentCube := strings.Fields(subset)
				cubesNumber, err := strconv.Atoi(currentCube[0])
				if err != nil {
					panic(err)
				}
				switch currentCube[1] {
				case "red":
					redCubes -= cubesNumber
				case "green":
					greenCubes -= cubesNumber
				case "blue":
					blueCubes -= cubesNumber
				default:
					panic("Invalid cube color")
				}
			}
			if redCubes < 0 || greenCubes < 0 || blueCubes < 0 {
				continue nextGame
			}
		}
		total += gameIndex + 1
	}
	return
}

func main() {
	puzzlePart := flag.Int("part", 1, "Puzzle part to solve (1 or 2)")
	flag.Parse()

	input := parsePuzzleInput("input.txt")
	switch *puzzlePart {
	case 1:
		fmt.Println(solvePartOne(input))
	case 2:
		fmt.Println(solvePartTwo(input))
	default:
		fmt.Println("Invalid puzzle part")
	}
}
