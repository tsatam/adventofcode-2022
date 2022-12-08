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

	scenicscores := findScenicScores(treemap)
	highestScore := highestScenicScore(scenicscores)

	fmt.Printf("pt1: %d\n", numVisibleTrees)
	fmt.Printf("pt2: %d\n", highestScore)
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

func findScenicScores(treemap [][]int) [][]int {

	scenicscore := make([][]int, len(treemap))

	for y, _ := range treemap {
		scenicscore[y] = make([]int, len(treemap[y]))
		for x, _ := range treemap[y] {
			if x == 0 || y == 0 || x == len(treemap[y])-1 || y == len(treemap[y])-1 {
				scenicscore[y][x] = 0
			} else {
				scenicscore[y][x] = checkScenicScore(treemap, x, y)
			}
		}
	}

	return scenicscore
}

func highestScenicScore(scenicscores [][]int) int {
	highest := 0
	for _, row := range scenicscores {
		for _, score := range row {
			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func checkScenicScore(treemap [][]int, x, y int) int {
	return checkScenicScoreLeft(treemap, x, y) *
		checkScenicScoreRight(treemap, x, y) *
		checkScenicScoreUp(treemap, x, y) *
		checkScenicScoreDown(treemap, x, y)
}

func checkScenicScoreLeft(treemap [][]int, x, y int) int {
	score := 0
	for xp := x - 1; xp >= 0; xp-- {
		score++
		if treemap[y][xp] >= treemap[y][x] {
			return score
		}
	}
	return score
}
func checkScenicScoreRight(treemap [][]int, x, y int) int {
	score := 0
	for xp := x + 1; xp < len(treemap[y]); xp++ {
		score++
		if treemap[y][xp] >= treemap[y][x] {
			return score
		}
	}
	return score
}
func checkScenicScoreUp(treemap [][]int, x, y int) int {
	score := 0
	for yp := y - 1; yp >= 0; yp-- {
		score++
		if treemap[yp][x] >= treemap[y][x] {
			return score
		}
	}
	return score
}
func checkScenicScoreDown(treemap [][]int, x, y int) int {
	score := 0
	for yp := y + 1; yp < len(treemap); yp++ {
		score++
		if treemap[yp][x] >= treemap[y][x] {
			return score
		}
	}
	return score
}
