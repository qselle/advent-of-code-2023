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

type Tile struct {
	Value  int
	Border bool
}

type Tiles [][]Tile

func isStartingPoint(pipe int) bool {
	return pipe&(North|South|East|West) == North|South|East|West
}

func (t Tiles) Move(y, x, count, from int) int {
	t[y][x].Border = true
	if isStartingPoint(t[y][x].Value) && from != North|South|East|West {
		return count
	}
	if from != South && t[y][x].Value&North != 0 && t[y-1][x].Value&South != 0 {
		return t.Move(y-1, x, count+1, North)
	}
	if from != North && t[y][x].Value&South != 0 && t[y+1][x].Value&North != 0 {
		return t.Move(y+1, x, count+1, South)
	}
	if from != West && t[y][x].Value&East != 0 && t[y][x+1].Value&West != 0 {
		return t.Move(y, x+1, count+1, East)
	}
	if from != East && t[y][x].Value&West != 0 && t[y][x-1].Value&East != 0 {
		return t.Move(y, x-1, count+1, West)
	}
	return 0
}

func (t Tiles) ComputeFarthestStep() int {
	for y, pipes := range t {
		for x, pipe := range pipes {
			if isStartingPoint(pipe.Value) {
				return t.Move(y, x, 0, North|South|East|West) / 2
			}
		}
	}
	return 0
}

func (t Tiles) Dump() {
	count := 0
	for _, pipes := range t {
		inside := false
		for _, pipe := range pipes {
			if pipe.Border {
				if isStartingPoint(pipe.Value) {
					fmt.Print("░")
					inside = !inside
				} else if pipe.Value&North != 0 && pipe.Value&South != 0 {
					fmt.Print("│")
					inside = !inside
				} else if pipe.Value&East != 0 && pipe.Value&West != 0 {
					fmt.Print("─")
				} else if pipe.Value&North != 0 && pipe.Value&East != 0 {
					fmt.Print("└")
					inside = !inside
				} else if pipe.Value&North != 0 && pipe.Value&West != 0 {
					fmt.Print("┘")
					inside = !inside
				} else if pipe.Value&South != 0 && pipe.Value&West != 0 {
					fmt.Print("┐")
					// inside = !inside
				} else if pipe.Value&South != 0 && pipe.Value&East != 0 {
					fmt.Print("┌")
					// inside = !inside
				}
			} else {
				if inside {
					count++
					fmt.Print("█")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("Part 2", count)
}
func parseTiles(input []string) Tiles {
	tiles := make([][]Tile, len(input))
	for y, line := range input {
		currentTiles := make([]Tile, len(line))
		for x, char := range line {
			char := string(char)
			currentTiles[x].Value = Connections[char]
		}
		tiles[y] = currentTiles
	}
	return tiles
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	tiles := parseTiles(input)
	fmt.Println("Part 1:", tiles.ComputeFarthestStep())
	tiles.Dump()
}
