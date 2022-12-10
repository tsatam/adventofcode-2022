package main

import (
	set "common"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input         string
	checkedCycles = set.NewSet(20, 60, 100, 140, 180, 220)
)

func main() {
	fmt.Printf("pt1: [%d]\n", simulate(input))
}

func simulate(input string) int {
	x := 1

	xAtCycle := map[int]int{}

	splitAddCommand := strings.ReplaceAll(strings.TrimSpace(input), " ", "\n")
	instructions := strings.Split(splitAddCommand, "\n")

	for i, instruction := range instructions {
		cycle := i + 1
		if checkedCycles.Contains(cycle) {
			xAtCycle[cycle] = x
		}

		if instruction != "noop" && instruction != "addx" {
			add, err := strconv.Atoi(instruction)
			if err != nil {
				log.Fatal(err)
			}
			x += add
		}
	}

	sum := 0
	for c, xAtC := range xAtCycle {
		sum += c * xAtC
	}
	return sum
}
