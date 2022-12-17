package main

import (
	"bytes"
	c "common/cartesian"
	"common/set"
)

const (
	MAX_CAVE_HEIGHT = 40
)

type Cave struct {
	chamber [][7]bool
	yoffset int

	jets []c.Direction

	jetIdx int
}

func (cave *Cave) Display() string {
	var buffer bytes.Buffer

	for y := len(cave.chamber) - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if cave.chamber[y][x] {
				buffer.WriteRune('#')
			} else {
				buffer.WriteRune('.')
			}
		}
		buffer.WriteRune('\n')
	}
	buffer.WriteString("=======\n")
	return buffer.String()
}

func (cave *Cave) dropRock(rock Rock) {
	startPoint := cave.getHighestPoint() + 3

	rockToMove := rock.AtOffset(startPoint)

	for ; true; cave.jetIdx = (cave.jetIdx + 1) % len(cave.jets) {
		d := cave.jets[cave.jetIdx]
		rockToMove, _ = cave.tryMoveRock(rockToMove, d)

		var success bool
		rockToMove, success = cave.tryMoveRock(rockToMove, c.Down)
		if !success {
			cave.jetIdx = (cave.jetIdx + 1) % len(cave.jets)
			break
		}
	}

	cave.placeRock(rockToMove)
}

func (cave *Cave) placeRock(rock Rock) {
	points := rock.points.Slice()

	for _, point := range points {
		cave.placePoint(point)
	}
}

func (cave *Cave) placePoint(p c.Point) {
	if -p.Y >= len(cave.chamber) {
		rowsToAdd := make([][7]bool, -p.Y-(len(cave.chamber)-1))
		cave.chamber = append(cave.chamber, rowsToAdd...)
	}

	cave.chamber[-p.Y][p.X] = true
}

func (cave *Cave) getHighestPoint() int {
	return len(cave.chamber)
}

func (cave *Cave) tryMoveRock(r Rock, d c.Direction) (Rock, bool) {
	newRock := Rock{set.NewSet[c.Point]()}
	for _, point := range r.points.Slice() {
		if newPoint, success := cave.tryMovePoint(point, d); success {
			newRock.points.Add(newPoint)
		} else {
			return r, false
		}
	}
	return newRock, true
}

func (cave *Cave) tryMovePoint(p c.Point, d c.Direction) (c.Point, bool) {
	newPoint := p.Move(d)

	if newPoint.X < 0 || newPoint.X > 6 || newPoint.Y > 0 || cave.isObstructed(newPoint) {
		return p, false
	}

	return newPoint, true
}

func (cave *Cave) isObstructed(p c.Point) bool {
	if -p.Y >= len(cave.chamber) {
		return false
	}

	return cave.chamber[-p.Y][p.X]
}

func (cave *Cave) getTopRow() byte {
	result := byte(0b00000000)
	if len(cave.chamber) == 0 {
		return result
	}

	rowToGet := len(cave.chamber) - 1

	for _, isOccupied := range cave.chamber[rowToGet] {
		result = result << 2
		if isOccupied {
			result += 0b00000001
		}
	}

	return result
}

func (cave *Cave) getTopograhy() (result [7]int) {
	y := len(cave.chamber) - 1

	for x := 0; x < 7; x++ {
		for yOffset := 0; yOffset <= y; yOffset++ {
			if cave.chamber[y-yOffset][x] {
				result[x] = yOffset
				break
			}
		}
	}

	return
}

/*
	10 rocks

	cycle len: 8, found at 1

	height of: 1rock + 8rock * 1 + 1rock
*/
