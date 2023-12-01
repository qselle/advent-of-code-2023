package main

import (
	"flag"
	"github.com/quentinselle/aoc/2023/day1"
)

func main() {
	puzzleDay := flag.Int("day", 1, "Puzzle day to solve (1-25)")
	puzzlePart := flag.Int("part", 1, "Puzzle part to solve (1 or 2)")
	puzzleInputPath := flag.String("input", "../input/day1.txt", "Path to puzzle input file")
	flag.Parse()

	switch *puzzleDay {
	case 1:
		day1.Day1(*puzzleInputPath, *puzzlePart)
	}
}
