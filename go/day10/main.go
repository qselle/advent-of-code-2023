package main

import (
	"fmt"

	"github.com/quentinselle/aoc/2023/go/utils"
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

func (t Tiles) SolveAndDumpPart2() int {
	count := 0
	for y, pipes := range t {
		// inside := false
		border := 0
		last := 0
		for x, pipe := range pipes {
			if pipe.Border {
				if isStartingPoint(pipe.Value) {
					start := 0
					fmt.Print("░")
					if y > 0 && t[y-1][x].Value&South != 0 {
						start = start | North
					}
					if x > 0 && t[y][x-1].Value&East != 0 {
						start = start | West
					}
					if y < len(t)-1 && t[y+1][x].Value&North != 0 {
						start = start | South
					}
					if x < len(pipes)-1 && t[y][x+1].Value&West != 0 {
						start = start | East
					}

					if start&North != 0 && start&South != 0 {
						border++
					} else if start&North != 0 && start&East != 0 {
						// in
						last = start
					} else if start&North != 0 && start&West != 0 {
						// out
						if last&South != 0 {
							border++
						}
					} else if start&South != 0 && start&West != 0 {
						// out
						if last&North != 0 {
							border++
						}
					} else if start&South != 0 && start&East != 0 {
						// in
						last = start
					}
				} else if pipe.Value&North != 0 && pipe.Value&South != 0 {
					fmt.Print("│")
					border++
				} else if pipe.Value&East != 0 && pipe.Value&West != 0 {
					fmt.Print("─")
				} else if pipe.Value&North != 0 && pipe.Value&East != 0 {
					// in
					fmt.Print("└")
					last = pipe.Value
				} else if pipe.Value&North != 0 && pipe.Value&West != 0 {
					// out
					fmt.Print("┘")
					if last&South != 0 {
						border++
					}
				} else if pipe.Value&South != 0 && pipe.Value&West != 0 {
					// out
					fmt.Print("┐")
					if last&North != 0 {
						border++
					}
				} else if pipe.Value&South != 0 && pipe.Value&East != 0 {
					// in
					fmt.Print("┌")
					last = pipe.Value
				}
			} else {
				if border%2 != 0 {
					fmt.Print("█")
					count++
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
	return count
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
	fmt.Println("Part 2:", tiles.SolveAndDumpPart2())
}
