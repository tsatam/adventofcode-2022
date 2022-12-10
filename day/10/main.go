package main

import (
	set "common/set"
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
	sum, screen := simulate(input)
	fmt.Printf("pt1: [%d]\n", sum)
	fmt.Printf("pt2: [%s]\n", print(screen))
}

func simulate(input string) (int, Screen) {
	x := 1

	var screen Screen

	xAtCycle := map[int]int{}

	splitAddCommand := strings.ReplaceAll(strings.TrimSpace(input), " ", "\n")
	instructions := strings.Split(splitAddCommand, "\n")

	for i, instruction := range instructions {
		cycle := i + 1
		screenPosition := iterToScreenPosition(i)

		if checkedCycles.Contains(cycle) {
			xAtCycle[cycle] = x
		}
		if x == screenPosition.X || x == screenPosition.X-1 || x == screenPosition.X+1 {
			screen[screenPosition.Y][screenPosition.X] = 1
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
	return sum, screen
}
