package main

import (
	c "common/cartesian"
	_ "embed"
	"fmt"
	"log"
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
	towerHeight = getTowerHeight(jets, 1000000000000)
	fmt.Printf("pt2: [%d]\n", towerHeight)
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

// used to store intermediate results to check for cycles
type Key struct {
	rockType   int
	jetIdx     int
	topography [7]int
}

type Value struct {
	rock   int
	topRow byte
	height int
}

func getTowerHeight(jets []c.Direction, rocks int) int {
	cave := Cave{
		chamber: make([][7]bool, 0),
		yoffset: 0,

		jets:   jets,
		jetIdx: 0,
	}

	return getTowerHeightInternal(jets, rocks, 0, cave)

}

func getTowerHeightInternal(jets []c.Direction, rocks int, rockOffset int, cave Cave) int {

	cycleCheck := make(map[Key]Value, 0)

	handledRecurrence := false

	for i := 0; i < rocks; i++ {
		rockType := (i + rockOffset) % 5

		currentHeight := cave.getHighestPoint() + cave.yoffset
		key := Key{
			rockType:   rockType,
			jetIdx:     cave.jetIdx,
			topography: cave.getTopograhy(),
		}
		value := Value{
			rock:   i,
			topRow: cave.getTopRow(),
			height: currentHeight,
		}

		existing, exists := cycleCheck[key]
		if !handledRecurrence {
			if exists {
				if existing.topRow != value.topRow {
					log.Fatal("bad key, chuck a debug statement here")
				} else {
					log.Default().Printf("Recurrence found at i_1=[%d], i_2=[%d]", existing.rock, value.rock)

					offset := existing.rock
					offsetHeight := existing.height

					cycleLength := value.rock - existing.rock
					cycleHeight := currentHeight - offsetHeight

					totalCycles := (rocks - offset) / cycleLength
					cyclesToAdd := totalCycles - 1
					cycleHeightToAdd := cyclesToAdd * cycleHeight

					cave.yoffset += cycleHeightToAdd // add remaining cycles
					i += cyclesToAdd * cycleLength   // set iterator to post-cycle remainder
					if i >= rocks {
						break
					}
					handledRecurrence = true
				}
			} else {
				cycleCheck[key] = value
			}
		}

		cave.dropRock(Rocks[rockType])

	}
	return cave.getHighestPoint() + cave.yoffset
}

func findHighestDistance(topography [7]int) int {
	highest := 0
	for _, val := range topography {
		if val > highest {
			highest = val
		}
	}
	return highest
}
