package main

import (
	set "common"
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
	motions := parseInput(input)

	fmt.Printf("pt1: [%d]\n", positionsVisited(motions, 2))
	fmt.Printf("pt2: [%d]\n", positionsVisited(motions, 10))
}

type Direction rune

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Right Direction = 'R'
	Left  Direction = 'L'
)

type Motion struct {
	Direction Direction
	Distance  int
}

func (m Motion) String() string {
	return fmt.Sprintf("[%c %d]", m.Direction, m.Distance)
}

type Point struct{ x, y int }

func (p Point) Move(direction Direction) Point {
	switch direction {
	case Up:
		return Point{p.x, p.y - 1}
	case Down:
		return Point{p.x, p.y + 1}
	case Left:
		return Point{p.x - 1, p.y}
	case Right:
		return Point{p.x + 1, p.y}
	default:
		return p
	}
}

func (p Point) String() string {
	return fmt.Sprintf("{%d,%d}", p.x, p.y)
}

func parseInput(input string) []Motion {
	split := strings.Split(strings.TrimSpace(input), "\n")

	motions := make([]Motion, len(split))

	for i, line := range split {
		var direction Direction
		var distance int
		if _, err := fmt.Sscanf(line, "%c %d", &direction, &distance); err != nil {
			log.Fatal(err)
		}

		motions[i] = Motion{direction, distance}
	}

	return motions
}

func positionsVisited(motions []Motion, ropeLength int) int {
	knots := make([]Point, ropeLength)

	for i := 0; i < ropeLength; i++ {
		knots[i] = Point{0, 0}
	}

	visited := set.NewSet(knots[ropeLength-1])

	for _, motion := range motions {
		for v := 0; v < motion.Distance; v++ {
			knots[0] = knots[0].Move(motion.Direction)

			for i := 1; i < ropeLength; i++ {
				knots[i] = reconcileTail(knots[i-1], knots[i])
			}
			visited.Add(knots[ropeLength-1])
		}
	}

	return visited.Size()
}

func reconcileTail(head Point, tail Point) Point {
	if head.x-tail.x >= -1 && head.x-tail.x <= 1 &&
		head.y-tail.y >= -1 && head.y-tail.y <= 1 {
		return tail
	}

	if head.x-tail.x < -1 {
		shift := tail.Move(Left)
		if head.y < tail.y {
			return shift.Move(Up)
		}
		if head.y > tail.y {
			return shift.Move(Down)
		}
		return shift
	}

	if head.x-tail.x > 1 {
		shift := tail.Move(Right)
		if head.y < tail.y {
			return shift.Move(Up)
		}
		if head.y > tail.y {
			return shift.Move(Down)
		}
		return shift
	}
	if head.y-tail.y < -1 {
		shift := tail.Move(Up)
		if head.x < tail.x {
			return shift.Move(Left)
		}
		if head.x > tail.x {
			return shift.Move(Right)
		}
		return shift
	}

	if head.y-tail.y > 1 {
		shift := tail.Move(Down)
		if head.x < tail.x {
			return shift.Move(Left)
		}
		if head.x > tail.x {
			return shift.Move(Right)
		}
		return shift
	}

	log.Fatalf("Head [%v], Tail [%v] too far to reconcile", head, tail)
	return tail
}
