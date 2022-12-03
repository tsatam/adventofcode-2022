package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("pt1: %d\n", processIndividually(input))
	fmt.Printf("pt2: %d\n", processInGroups(input))
}

func processIndividually(input string) int {
	sum := 0

	for _, rucksack := range strings.Split(input, "\n") {
		sum += processRucksack(rucksack)
	}

	return sum
}

func processInGroups(input string) int {
	sum := 0
	rucksacks := strings.Split(input, "\n")

	for i := 0; i < len(rucksacks)-2; i += 3 {
		sum += findCommon([]string{rucksacks[i], rucksacks[i+1], rucksacks[i+2]})
	}
	return sum
}

func processRucksack(rucksack string) int {
	size := len(rucksack) / 2

	first, second := rucksack[:size], rucksack[size:]

	return findCommon([]string{first, second})
}

func findCommon(rucksacks []string) int {
	for idx, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		allContains := true
		for _, rucksack := range rucksacks {
			if !strings.ContainsRune(rucksack, char) {
				allContains = false
				break
			}
		}

		if allContains == true {
			return idx + 1
		}
	}

	return 0
}
