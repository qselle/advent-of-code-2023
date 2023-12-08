package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/quentinselle/aoc/2023/utils"
)

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Occurences map[string]int

func (o Occurences) GetOccurencesType() Type {
	if len(o) == 1 {
		return FiveOfAKind
	}
	for _, occurrence := range o {
		if occurrence == 4 {
			return FourOfAKind
		}
	}
	if len(o) == 2 {
		return FullHouse
	}
	for _, occurrence := range o {
		if occurrence == 3 {
			return ThreeOfAKind
		}
	}
	if len(o) == 3 {
		return TwoPair
	}
	for _, occurrence := range o {
		if occurrence == 2 {
			return OnePair
		}
	}
	return HighCard

}

var Strength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

type Card struct {
	Value    string
	Strength int
}

type Hand struct {
	Cards []Card
	Bid   int
	Type  Type
}

type Hands []Hand

func parseHands(input []string) Hands {
	hands := make(Hands, 0)
	for _, line := range input {
		cards := strings.Fields(line)
		bid, err := strconv.Atoi(cards[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{Bid: bid}
		occurrences := make(Occurences)
		for _, card := range cards[0] {
			card := string(card)
			hand.Cards = append(hand.Cards, Card{
				Value:    card,
				Strength: Strength[card],
			})
			if occurrence, ok := occurrences[card]; ok {
				occurrences[card] = occurrence + 1
			} else {
				occurrences[card] = 1
			}
		}
		hand.Type = occurrences.GetOccurencesType()
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	input := utils.ReadFileByLine("day7/input.txt")
	fmt.Println(parseHands(input))
}
