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
	treemap := parseInput(input)
	visibletrees := findVisibleTrees(treemap)
	numVisibleTrees := sumOfVisibleTrees(visibletrees)

	fmt.Printf("pt1: %d\n", numVisibleTrees)
}

func parseInput(input string) [][]int {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	treemap := make([][]int, len(rows))

	for y, row := range rows {
		treemap[y] = make([]int, len(row))
		for x, col := range strings.Split(row, "") {
			if val, err := strconv.Atoi(col); err != nil {
				log.Fatal(err)
			} else {
				treemap[y][x] = val
			}
		}
	}

	return treemap
}

func findVisibleTrees(treemap [][]int) [][]int {
	visibletrees := make([][]int, len(treemap))

	for y, _ := range treemap {
		visibletrees[y] = make([]int, len(treemap[y]))
		for x, _ := range treemap[y] {
			if x == 0 || y == 0 || x == len(treemap[y])-1 || y == len(treemap[y])-1 || checkVisibility(treemap, x, y) {
				visibletrees[y][x] = 1
			} else {
				visibletrees[y][x] = 0
			}
		}
	}

	return visibletrees
}

func sumOfVisibleTrees(visibletrees [][]int) int {
	sum := 0
	for _, row := range visibletrees {
		for _, isVisible := range row {
			sum += isVisible
		}
	}
	return sum
}

func checkVisibility(treemap [][]int, x, y int) bool {
	return checkVisibilityLeft(treemap, x, y) ||
		checkVisibilityRight(treemap, x, y) ||
		checkVisibilityUp(treemap, x, y) ||
		checkVisibilityDown(treemap, x, y)
}

func checkVisibilityLeft(treemap [][]int, x, y int) bool {
	for xp := 0; xp < x; xp++ {
		if treemap[y][xp] >= treemap[y][x] {
			return false
		}
	}
	return true
}
func checkVisibilityRight(treemap [][]int, x, y int) bool {
	for xp := len(treemap[y]) - 1; xp > x; xp-- {
		if treemap[y][xp] >= treemap[y][x] {
			return false
		}
	}
	return true
}
func checkVisibilityUp(treemap [][]int, x, y int) bool {
	for yp := 0; yp < y; yp++ {
		if treemap[yp][x] >= treemap[y][x] {
			return false
		}
	}
	return true
}
func checkVisibilityDown(treemap [][]int, x, y int) bool {
	for yp := len(treemap) - 1; yp > y; yp-- {
		if treemap[yp][x] >= treemap[y][x] {
			return false
		}
	}
	return true
}
