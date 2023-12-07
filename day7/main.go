package main

import (
	"fmt"
	"github.com/quentinselle/aoc/2023/utils"
	"strconv"
	"strings"
)

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HigCard
)

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
	Type  int
}

func (h Hand) GetHandType() {
	previous := ""
	count := 0
	for _, card := range h.Cards {
		if previous == "" {
			previous = card.Value
		}
		if card.Value == previous {
			count++
		}
	}
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
		for _, card := range cards[0] {
			card := string(card)
			hand.Cards = append(hand.Cards, Card{
				Value:    card,
				Strength: Strength[card],
			})
		}
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	input := utils.ReadFileByLine("day7/input.txt")
	fmt.Println(parseHands(input))
}
