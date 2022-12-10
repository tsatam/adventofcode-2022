package main

import (
	cartesian "common/cartesian"
	set "common/set"
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

type Motion struct {
	Direction cartesian.Direction
	Distance  int
}

func (m Motion) String() string {
	return fmt.Sprintf("[%c %d]", m.Direction, m.Distance)
}

func parseInput(input string) []Motion {
	split := strings.Split(strings.TrimSpace(input), "\n")

	motions := make([]Motion, len(split))

	for i, line := range split {
		var direction cartesian.Direction
		var distance int
		if _, err := fmt.Sscanf(line, "%c %d", &direction, &distance); err != nil {
			log.Fatal(err)
		}

		motions[i] = Motion{direction, distance}
	}

	return motions
}

func positionsVisited(motions []Motion, ropeLength int) int {
	knots := make([]cartesian.Point, ropeLength)

	for i := 0; i < ropeLength; i++ {
		knots[i] = cartesian.Point{X: 0, Y: 0}
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

func reconcileTail(head cartesian.Point, tail cartesian.Point) cartesian.Point {
	if head.X-tail.X >= -1 && head.X-tail.X <= 1 &&
		head.Y-tail.Y >= -1 && head.Y-tail.Y <= 1 {
		return tail
	}

	if head.X-tail.X < -1 {
		shift := tail.Move(cartesian.Left)
		if head.Y < tail.Y {
			return shift.Move(cartesian.Up)
		}
		if head.Y > tail.Y {
			return shift.Move(cartesian.Down)
		}
		return shift
	}

	if head.X-tail.X > 1 {
		shift := tail.Move(cartesian.Right)
		if head.Y < tail.Y {
			return shift.Move(cartesian.Up)
		}
		if head.Y > tail.Y {
			return shift.Move(cartesian.Down)
		}
		return shift
	}
	if head.Y-tail.Y < -1 {
		shift := tail.Move(cartesian.Up)
		if head.X < tail.X {
			return shift.Move(cartesian.Left)
		}
		if head.X > tail.X {
			return shift.Move(cartesian.Right)
		}
		return shift
	}

	if head.Y-tail.Y > 1 {
		shift := tail.Move(cartesian.Down)
		if head.X < tail.X {
			return shift.Move(cartesian.Left)
		}
		if head.X > tail.X {
			return shift.Move(cartesian.Right)
		}
		return shift
	}

	log.Fatalf("Head [%v], Tail [%v] too far to reconcile", head, tail)
	return tail
}
