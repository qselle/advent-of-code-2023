package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/quentinselle/aoc/2023/go/utils"
)

type ScratchCard struct {
	WinningNumbers []int
	CardNumbers    []int
}

func (sc ScratchCard) TotalMatches() int {
	total := 0
	for _, numbers := range sc.CardNumbers {
		if slices.Contains(sc.WinningNumbers, numbers) {
			total++
		}
	}
	return total
}

func (sc ScratchCard) TotalPointsWorth() int {
	total := 0
	for index := 0; index < sc.TotalMatches(); index++ {
		if index == 0 {
			total = 1
		} else {
			total *= 2
		}
	}
	return total
}

type ScratchCards []ScratchCard

func (sc ScratchCards) TotalPointsWorth() int {
	total := 0
	for _, scratchCard := range sc {
		total += scratchCard.TotalPointsWorth()
	}
	return total
}

func (sc ScratchCards) TotalScratchCardsWin(start, end, scNumbers int) int {
	for index := start; index < end; index++ {
		matches := sc[index].TotalMatches()
		if matches > 0 {
			// assign index to the next scratch card (padding)
			index := index + 1
			scNumbers = sc.TotalScratchCardsWin(index, index+matches, scNumbers+matches)
		}
	}
	return scNumbers
}

func parseScratchCards(input []string) ScratchCards {
	scratchCards := make(ScratchCards, 0)

	for _, line := range input {
		game := strings.Split(line, ":")[1]
		allNumbers := strings.Split(game, "|")
		scratchCards = append(scratchCards, ScratchCard{
			utils.ExtractNumbersFromString(allNumbers[0]),
			utils.ExtractNumbersFromString(allNumbers[1]),
		})
	}
	return scratchCards
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	sc := parseScratchCards(input)
	fmt.Println("Part 1:", sc.TotalPointsWorth())
	fmt.Println("Part 2:", sc.TotalScratchCardsWin(0, len(sc), len(sc)))
}
