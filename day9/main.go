package main

import (
	"fmt"
	"slices"

	"github.com/quentinselle/aoc/2023/utils"
)

type Sequence []int

func (s Sequence) Predict() Sequence {
	prev := s[0]
	predictions := Sequence{}
	for _, current := range s[1:] {
		currentPrediction := current - prev
		prev = current
		predictions = append(predictions, currentPrediction)
	}
	return predictions
}

type History []int

func (h History) Predict() []Sequence {
	predictions := append([]Sequence{}, Sequence(h))
	for index := 0; slices.Max(predictions[index]) != 0 || slices.Min(predictions[index]) != 0; index++ {
		predictions = append(predictions, predictions[index].Predict())
	}
	return predictions
}

func (h History) ExtrapolateBackwards() int {
	sequences := h.Predict()
	last := 0
	for i := len(sequences) - 1; i > 0; i-- {
		last = sequences[i-1][0] - last
	}
	return last
}

func (h History) Extrapolate() int {
	sequences := h.Predict()
	extrapolate := 0
	for i := len(sequences) - 1; i > 0; i-- {
		extrapolate = sequences[i-1][len(sequences[i-1])-1] + extrapolate
	}
	return extrapolate
}

type Report struct {
	Histories []History
}

func (r Report) ExtrapolateBackwards() int {
	sum := 0
	for _, history := range r.Histories {
		sum += history.ExtrapolateBackwards()
	}
	return sum
}

func (r Report) Extrapolate() int {
	sum := 0
	for _, history := range r.Histories {
		sum += history.Extrapolate()
	}
	return sum
}

func parseHistories(input []string) Report {
	report := Report{}
	for _, line := range input {
		report.Histories = append(report.Histories, utils.ExtractNumbersFromString(line))
	}
	return report
}

func main() {
	input := utils.ReadFileByLine("input.txt")
	report := parseHistories(input)
	fmt.Println("Part 1:", report.Extrapolate())
	fmt.Println("Part 2:", report.ExtrapolateBackwards())
}
