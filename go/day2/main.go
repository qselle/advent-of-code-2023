package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quentinselle/aoc/2023/go/utils"
)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []Set
}

func parsePuzzleInput(input []string) []Game {
	var games []Game
	var err error

	for _, line := range input {
		currentGame := Game{}
		byGameId := strings.Split(line, ":")
		currentGame.id, err = strconv.Atoi(strings.TrimPrefix(byGameId[0], "Game "))
		if err != nil {
			panic(err)
		}
		for _, set := range strings.Split(byGameId[1], ";") {
			currentSet := Set{}

			for _, subset := range strings.Split(set, ",") {
				currentCube := strings.Fields(subset)
				cubesNumber, err := strconv.Atoi(currentCube[0])
				if err != nil {
					panic(err)
				}

				switch currentCube[1] {
				case "red":
					currentSet.red += cubesNumber
				case "green":
					currentSet.green += cubesNumber
				case "blue":
					currentSet.blue += cubesNumber
				default:
					panic("Invalid cube color")
				}
			}
			currentGame.sets = append(currentGame.sets, currentSet)
		}
		games = append(games, currentGame)
	}
	return games
}

func solvePartTwo(games []Game) int {
	sum := 0

	for _, game := range games {
		redCubesMin, greenCubesMin, blueCubesMin := 0, 0, 0
		for _, gameSet := range game.sets {
			if gameSet.red > redCubesMin {
				redCubesMin = gameSet.red
			}
			if gameSet.green > greenCubesMin {
				greenCubesMin = gameSet.green
			}
			if gameSet.blue > blueCubesMin {
				blueCubesMin = gameSet.blue
			}
		}
		sum += redCubesMin * greenCubesMin * blueCubesMin
	}
	return sum
}

func solvePartOne(games []Game) int {
	const redCubesInBag = 12
	const greenCubesInBag = 13
	const blueCubesInBag = 14
	sum := 0

nextGame:
	for _, game := range games {
		for _, gameSet := range game.sets {
			if gameSet.red > redCubesInBag ||
				gameSet.green > greenCubesInBag ||
				gameSet.blue > blueCubesInBag {
				continue nextGame
			}
		}
		sum += game.id
	}
	return sum
}

func main() {
	input := parsePuzzleInput(utils.ReadFileByLine("input.txt"))
	fmt.Println("Part 1:", solvePartOne(input))
	fmt.Println("Part 2:", solvePartTwo(input))
}
