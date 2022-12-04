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
	fmt.Printf("pt1: %v\n", countFullContains(input))
	fmt.Printf("pt2: %v\n", countOverlaps(input))
}

func countFullContains(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			if hasFullContain(line) {
				sum++
			}
		}
	}
	return sum
}

func countOverlaps(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			if hasOverlap(line) {
				sum++
			}
		}
	}
	return sum
}

func hasFullContain(line string) bool {
	first, second := parseLine(line)

	return first.contains(second) || second.contains(first)
}

func hasOverlap(line string) bool {
	first, second := parseLine(line)

	return first.overlaps(second)
}

type Assignment struct {
	start int
	end   int
}

func parseLine(line string) (Assignment, Assignment) {
	splitLine := strings.Split(line, ",")
	return parseAssignment(splitLine[0]), parseAssignment(splitLine[1])
}

func parseAssignment(raw string) Assignment {
	split := strings.Split(raw, "-")
	start, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	return Assignment{start, end}
}

func (a *Assignment) contains(b Assignment) bool {
	return a.start <= b.start && a.end >= b.end
}

func (a *Assignment) overlaps(b Assignment) bool {
	if a.end < b.start || b.end < a.start {
		return false
	} else {
		return true
	}
}
