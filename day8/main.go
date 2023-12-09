package main

import (
	"fmt"
	"strings"

	"github.com/quentinselle/aoc/2023/utils"
)

type Node struct {
	Left  string
	Right string
}

type Maps struct {
	Directions []string
	Nodes      map[string]Node
}

func (m Maps) StepsToReachEnd() int {
	nodeName := "AAA"
	steps := 0
	for {
		for _, direction := range m.Directions {
			if direction == "R" {
				nodeName = m.Nodes[nodeName].Right
			} else {
				nodeName = m.Nodes[nodeName].Left
			}
			steps++
			if nodeName == "ZZZ" {
				return steps
			}
		}

	}
}

func parseMaps(input []string) Maps {
	maps := Maps{}

	maps.Directions = make([]string, len(input[0]))
	for i, direction := range input[0] {
		maps.Directions[i] = string(direction)
	}
	maps.Nodes = make(map[string]Node)
	for _, line := range input[2:] {
		node := Node{}
		parts := strings.Split(line, " = ")
		str := strings.Split(strings.Trim(parts[1], "()"), ", ")
		node.Left = str[0]
		node.Right = str[1]
		maps.Nodes[parts[0]] = node
	}
	return maps
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	maps := (parseMaps(input))
	fmt.Println(maps.StepsToReachEnd())
}
