package main

import (
	c "common/cartesian"
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	jets := parseInput(input)
	towerHeight := getTowerHeight(jets, 2022)

	fmt.Printf("pt1: [%d]\n", towerHeight)
}

func parseInput(input string) []c.Direction {
	result := make([]c.Direction, len(input))
	for i, d := range strings.Split(input, "") {
		if d == "<" {
			result[i] = c.Left
		} else if d == ">" {
			result[i] = c.Right
		}
	}

	return result
}

func getTowerHeight(jets []c.Direction, rocks int) int {
	cave := Cave{
		chamber: make([][7]bool, 0),
		yoffset: 0,

		jets:   jets,
		jetIdx: 0,
	}

	for i := 0; i < rocks; i++ {
		cave.dropRock(Rocks[i%5])
	}
	return cave.getHighestPoint() + cave.yoffset
}
