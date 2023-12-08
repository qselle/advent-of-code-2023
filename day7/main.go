package main

import (
	"fmt"
	"sort"
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

	MaxCardByHand = 5
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

func (o Occurences) GetOccurencesTypeJoker() Type {
	if len(o) == 1 ||
		(len(o) == 2 && o["J"] > 0) {
		return FiveOfAKind
	}
	for label, occurrence := range o {
		if occurrence == 4 ||
			(occurrence == 3 && o["J"] > 0) ||
			(label != "J" && occurrence == 2 && o["J"] == 2) {
			return FourOfAKind
		}
	}
	if len(o) == 2 ||
		(len(o) == 3 && o["J"] > 0) {
		return FullHouse
	}
	for _, occurrence := range o {
		if occurrence == 3 ||
			(occurrence == 2 && o["J"] > 0) {
			return ThreeOfAKind
		}
	}
	if len(o) == 3 ||
		(len(o) == 4 && o["J"] > 0) {
		return TwoPair
	}
	for _, occurrence := range o {
		if occurrence == 2 || o["J"] > 0 {
			return OnePair
		}
	}
	return HighCard

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

func (h Hands) SortDesc() {
	sort.Slice(h, func(i, j int) bool {
		if h[i].Type < h[j].Type {
			return true
		}
		if h[i].Type == h[j].Type {
			for k := 0; k < MaxCardByHand; k++ {
				if h[i].Cards[k].Strength == h[j].Cards[k].Strength {
					continue
				}
				if h[i].Cards[k].Strength < h[j].Cards[k].Strength {
					return true
				}
				return false
			}
		}
		return false
	})
}

func (h Hands) ComputeTotalWinnings() int {
	total := 0
	for i, hand := range h {
		total += hand.Bid * (i + 1)
	}
	return total
}

func parseHands(input []string) Hands {
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

func parseHandsJoker(input []string) Hands {
	var StrengthJoker = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

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
				Strength: StrengthJoker[card],
			})
			if occurrence, ok := occurrences[card]; ok {
				occurrences[card] = occurrence + 1
			} else {
				occurrences[card] = 1
			}
		}
		hand.Type = occurrences.GetOccurencesTypeJoker()
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	hands := parseHands(input)
	hands.SortDesc()
	fmt.Println("Part 1:", hands.ComputeTotalWinnings())
	hands = parseHandsJoker(input)
	hands.SortDesc()
	fmt.Println("Part 2:", hands.ComputeTotalWinnings())
}
