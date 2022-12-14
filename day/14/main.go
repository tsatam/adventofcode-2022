package main

import (
	c "common/cartesian"
	s "common/set"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

var (
	//go:embed input
	input      string
	sandSource = c.Point{X: 500, Y: 0}
)

func main() {
	lines := parseInput(input)
	rocks := getPointsForLines(lines)

	numSand := processSand(rocks, false)
	numSandWithFloor := processSand(rocks, true)

	fmt.Printf("pt1: [%d]\n", numSand)
	fmt.Printf("pt2: [%d]\n", numSandWithFloor)

}

type Line = []c.Point

func parseInput(input string) []Line {
	rawLines := strings.Split(strings.TrimSpace(input), "\n")

	lines := make([]Line, len(rawLines))

	for i, line := range rawLines {
		rawPoints := strings.Split(line, " -> ")
		lines[i] = make(Line, len(rawPoints))
		for j, rawPoint := range rawPoints {
			point := parsePoint(rawPoint)
			lines[i][j] = point
		}
	}
	return lines
}

func parsePoint(raw string) c.Point {
	var x, y int

	if _, err := fmt.Sscanf(raw, "%d,%d", &x, &y); err != nil {
		log.Fatal(err)
	}

	return c.Point{X: x, Y: y}
}

func getPointsForLines(lines []Line) s.Set[c.Point] {
	points := s.NewSet[c.Point]()

	for _, line := range lines {
		points.Merge(getPointsForLine(line))
	}

	return points
}

func getPointsForLine(line Line) s.Set[c.Point] {
	points := s.NewSet[c.Point]()
	for i := 0; i < len(line)-1; i++ {
		points.Merge(getPointsBetweenPoints(line[i], line[i+1]))
	}

	return points
}

func getPointsBetweenPoints(a, b c.Point) s.Set[c.Point] {
	points := s.NewSet(a, b)

	switch a.IsInDirection(b) {
	case c.Up:
		for y := a.Y - 1; y > b.Y; y-- {
			points.Add(c.Point{X: a.X, Y: y})
		}

	case c.Down:
		for y := a.Y + 1; y < b.Y; y++ {
			points.Add(c.Point{X: a.X, Y: y})
		}
	case c.Left:
		for x := a.X - 1; x > b.X; x-- {
			points.Add(c.Point{X: x, Y: a.Y})
		}

	case c.Right:
		for x := a.X + 1; x < b.X; x++ {
			points.Add(c.Point{X: x, Y: a.Y})
		}
	}

	return points
}

func processSand(rocks s.Set[c.Point], shouldFloor bool) int {
	rocksAndSand := s.NewSet[c.Point]()
	rocksAndSand.Merge(rocks)

	lowestY := findLowestY(rocks)
	floor := lowestY + 2

	settledSand := 0
	oops := false

	for !oops {
		sand := sandSource
		sandHasSettled := false

		for !sandHasSettled {

			if !shouldFloor && sand.Y > lowestY {
				oops = true
				break
			}

			if shouldFloor && sand.Y == floor-1 {
				sandHasSettled = true
				rocksAndSand.Add(sand)
				settledSand++
			}

			if newSand := sand.Move(c.Down); !rocksAndSand.Contains(newSand) {
				sand = newSand
				continue
			} else if newSand := sand.Move(c.Down).Move(c.Left); !rocksAndSand.Contains(newSand) {
				sand = newSand
				continue
			} else if newSand := sand.Move(c.Down).Move(c.Right); !rocksAndSand.Contains(newSand) {
				sand = newSand
			} else {
				sandHasSettled = true
				rocksAndSand.Add(sand)
				settledSand++
			}
		}

		if sandHasSettled && sand == sandSource {
			oops = true
		}
	}

	return settledSand
}

// "lowest" means highest Y value
func findLowestY(rocks s.Set[c.Point]) int {
	lowestY := 0

	for _, rock := range rocks.Slice() {
		if rock.Y > lowestY {
			lowestY = rock.Y
		}
	}

	return lowestY
}
