package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("%d\n", readInput(input))
}

func readInput(input string) int {
	lines := strings.Split(input, "\n")

	highestSum := 0
	currentSum := 0

	for _, line := range lines {
		if line == "" {
			if currentSum > highestSum {
				highestSum = currentSum
			}
			currentSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentSum += calories
		}
	}

	return highestSum
}
