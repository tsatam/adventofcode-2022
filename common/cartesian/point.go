package cartesian

import "log"

type Point struct{ X, Y int }

func (p Point) Move(direction Direction) Point {
	switch direction {
	case Up:
		return Point{p.X, p.Y - 1}
	case Down:
		return Point{p.X, p.Y + 1}
	case Left:
		return Point{p.X - 1, p.Y}
	case Right:
		return Point{p.X + 1, p.Y}
	default:
		return p
	}
}

func (p Point) IsInDirection(other Point) Direction {
	if p.X == other.X {
		if p.Y < other.Y {
			return Down
		}
		if p.Y > other.Y {
			return Up
		}
	}
	if p.Y == other.Y {
		if p.X < other.X {
			return Right
		}
		if p.X > other.X {
			return Left
		}
	}

	log.Fatalf("Points [%v,%v] do not share axis", p, other)
	return Down
}

func (p Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
