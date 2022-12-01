package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("%d\n", threeHighestSums(input))
}

func highestSum(input string) int {
	sums := readInput(input)
	sort.Ints(sums)
	return sums[len(sums)-1]
}

func threeHighestSums(input string) int {
	sums := readInput(input)
	sort.Ints(sums)
	length := len(sums)
	return sums[length-1] + sums[length-2] + sums[length-3]
}

func readInput(input string) []int {
	lines := strings.Split(input, "\n")

	sums := []int{}
	currentSum := 0

	for _, line := range lines {
		if line == "" {
			sums = append(sums, currentSum)
			currentSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentSum += calories
		}
	}

	return sums
}
