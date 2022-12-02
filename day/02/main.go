package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input
	input       string
	choiceRanks = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
)

func main() {
	pt1, pt2 := processGame(input)
	fmt.Printf("pt1: %d\npt2: %d\n", pt1, pt2)
}

func processGame(input string) (int, int) {
	pt1sum, pt2sum := 0, 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 3 {
			pt1sum += processLinePt1(line)
			pt2sum += processLinePt2(line)
		}
	}
	return pt1sum, pt2sum
}

func processLinePt1(line string) int {
	choices := strings.Split(line, " ")
	opponentChoice := choiceRanks[choices[0]]
	playerChoice := choiceRanks[choices[1]]

	return calculateScoreForLine(opponentChoice, playerChoice)
}

func processLinePt2(line string) int {
	choices := strings.Split(line, " ")
	opponentChoice := choiceRanks[choices[0]]
	desiredResult := choiceRanks[choices[1]] - 1
	playerChoice := (opponentChoice + desiredResult) % 3

	if playerChoice == -1 {
		playerChoice = 2
	} // go modulo preserves negative numbers so this needs to be explicitly handled

	return calculateScoreForLine(opponentChoice, playerChoice)
}

func calculateScoreForLine(opponentChoice int, playerChoice int) int {
	choiceScore := playerChoice + 1
	resultScore := 0

	if opponentChoice == 0 && playerChoice == 2 {
		resultScore = 0
	} else if (opponentChoice == 2 && playerChoice == 0) || opponentChoice < playerChoice {
		resultScore = 6
	} else if opponentChoice == playerChoice {
		resultScore = 3
	} else {
		resultScore = 0
	}

	return choiceScore + resultScore
}
