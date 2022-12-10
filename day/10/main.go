package main

import (
	cartesian "common/cartesian"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	sum, screen := simulate(input)
	fmt.Printf("pt1: [%d]\n", sum)
	fmt.Printf("pt2: [%s]\n", print(screen))
}

func simulate(input string) (int, Screen) {
	x := 1

	sum := 0
	var screen Screen

	for i, instruction := range getInstructions(input) {
		if i == 19 || i == 59 || i == 99 || i == 139 || i == 179 || i == 219 {
			sum += x * (i + 1)
		}
		screenPosition := cartesian.Point{X: i % 40, Y: i / 40}
		if x == screenPosition.X || x == screenPosition.X-1 || x == screenPosition.X+1 {
			screen[screenPosition.Y][screenPosition.X] = 1
		}

		if instruction != "noop" && instruction != "addx" {
			add, _ := strconv.Atoi(instruction)
			x += add
		}
	}

	return sum, screen
}

func getInstructions(input string) []string {
	splitAddCommand := strings.ReplaceAll(strings.TrimSpace(input), " ", "\n")
	return strings.Split(splitAddCommand, "\n")
}
