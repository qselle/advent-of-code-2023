package main

import (
	"fmt"

	"github.com/quentinselle/aoc/2023/utils"
)

const (
	Ground = 0
	North  = 1 << iota
	South
	East
	West
)

var (
	Connections = map[string]int{
		"|": North | South,
		"-": East | West,
		"L": North | East,
		"J": North | West,
		"7": South | West,
		"F": South | East,
		"S": North | South | East | West,
		".": Ground,
	}
)

type Tiles [][]int

func isStartingPoint(pipe int) bool {
	return pipe&(North|South|East|West) == North|South|East|West
}

func (t Tiles) Move(y, x, count, from int) int {
	if isStartingPoint(t[y][x]) && count > 0 {
		return count
	}
	if from != South && t[y][x]&North != 0 && t[y-1][x]&South != 0 {
		return t.Move(y-1, x, count+1, North)
	}
	if from != North && t[y][x]&South != 0 && t[y+1][x]&North != 0 {
		return t.Move(y+1, x, count+1, South)
	}
	if from != West && t[y][x]&East != 0 && t[y][x+1]&West != 0 {
		return t.Move(y, x+1, count+1, East)
	}
	if from != East && t[y][x]&West != 0 && t[y][x-1]&East != 0 {
		return t.Move(y, x-1, count+1, West)
	}
	return 0
}

func (t Tiles) ComputeFarthestStep() int {
	for y, pipes := range t {
		for x, pipe := range pipes {
			if isStartingPoint(pipe) {
				return t.Move(y, x, 0, North|South|East|West) / 2
			}
		}
	}
	return 0
}

func parseTiles(input []string) Tiles {
	tiles := make([][]int, len(input))
	for y, line := range input {
		currentTiles := make([]int, len(line))
		for x, char := range line {
			char := string(char)
			currentTiles[x] = Connections[char]
		}
		tiles[y] = currentTiles
	}
	return tiles
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	tiles := parseTiles(input)
	fmt.Println("Part 1:", tiles.ComputeFarthestStep())
	fmt.Println("Part 2:")
}
