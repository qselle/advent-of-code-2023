package main

import (
	"fmt"
	"github.com/quentinselle/aoc/2023/utils"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	WinningNumbers []int
	CardNumbers    []int
}

func (sc ScratchCard) TotalPointsWorth() int {
	total := 0
	for _, numbers := range sc.CardNumbers {
		if slices.Contains(sc.WinningNumbers, numbers) {
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}
	return total
}

type ScratchCards map[int]ScratchCard

func (sc ScratchCards) TotalPointsWorth() int {
	total := 0
	for _, scratchCard := range sc {
		total += scratchCard.TotalPointsWorth()
	}
	return total
}

func extractNumbersFromString(str string) []int {
	numbers := make([]int, 0)

	for _, winningNumbers := range strings.Fields(str) {
		number, err := strconv.Atoi(winningNumbers)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func parseScratchCards(input []string) ScratchCards {
	scratchCards := make(ScratchCards)

	for _, line := range input {
		byCardId := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(byCardId[0], "Card ")))
		if err != nil {
			panic(err)
		}
		allNumbers := strings.Split(byCardId[1], "|")
		scratchCards[id] = ScratchCard{
			extractNumbersFromString(allNumbers[0]),
			extractNumbersFromString(allNumbers[1]),
		}
	}
	return scratchCards
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	sc := parseScratchCards(input)
	fmt.Println("Part 1:", sc.TotalPointsWorth())
}
