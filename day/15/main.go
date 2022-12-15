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
	input string
)

func main() {
	sensors := parseInput(input)

	countOccupiedAt := getCountOccupiedAt(sensors, 2000000)
	fmt.Printf("pt1: [%d]\n", countOccupiedAt)

	tuningFrequency := getDistressBeaconTuningFrequency(sensors, 4000000)
	fmt.Printf("pt2: [%d]\n", tuningFrequency)
}

type Sensor struct {
	position      c.Point
	closestBeacon c.Point

	distance int
}

func parseInput(input string) []Sensor {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sensors := make([]Sensor, len(lines))

	for i, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int

		if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY); err != nil {
			log.Fatal(err)
		}

		position := c.Point{X: sensorX, Y: sensorY}
		closestBeacon := c.Point{X: beaconX, Y: beaconY}
		sensors[i] = Sensor{
			position:      position,
			closestBeacon: closestBeacon,
			distance:      position.ManhattanDistance(closestBeacon),
		}
	}

	return sensors
}

func getOccupiedPositions(sensors []Sensor, y int) s.Set[c.Point] {
	occupied := s.NewSet[c.Point]()

	for _, sensor := range sensors {
		occupied.Union(sensor.getOccupiedAtRow(y))
	}
	for _, sensor := range sensors {
		if sensor.closestBeacon.Y == y {
			occupied.Remove(sensor.closestBeacon)
		}
	}
	return occupied
}

func getCountOccupiedAt(sensors []Sensor, y int) int {
	occupied := getOccupiedPositions(sensors, y)
	return occupied.Size()
}

func getDistressBeaconTuningFrequency(sensors []Sensor, bounds int) int {
	pointsToCheck := getIntersectionPoints(sensors, bounds)

	for _, point := range pointsToCheck.Slice() {
		invalid := false

		for _, sensor := range sensors {
			if sensor.isInRange(point) {
				invalid = true
				break
			}
		}

		if !invalid {
			return point.X*4000000 + point.Y
		}
	}
	return -1
}

func getIntersectionPoints(sensors []Sensor, bounds int) s.Set[c.Point] {
	points := s.NewSet(
		c.Point{X: 0, Y: 0},
		c.Point{X: bounds, Y: 0},
		c.Point{X: 0, Y: bounds},
		c.Point{X: bounds, Y: bounds},
	)

	ascLines := make([]Line, 0)
	descLines := make([]Line, 0)

	for _, sensor := range sensors {
		nw, ne, se, sw := sensor.getBoundaries()

		ascLines = append(ascLines, nw, se)
		descLines = append(descLines, ne, sw)
	}

	for _, ascLine := range ascLines {
		for _, descLine := range descLines {
			p := intersects(ascLine, descLine)

			if p.X >= 0 && p.X <= bounds && p.Y >= 0 && p.Y <= bounds {
				points.Add(p)
			}
		}
	}

	return points
}

func (sensor *Sensor) isInRange(p c.Point) bool {
	return sensor.position.ManhattanDistance(p) <= sensor.position.ManhattanDistance(sensor.closestBeacon)
}

func (sensor *Sensor) getOccupiedAtRow(y int) s.Set[c.Point] {
	occupied := s.NewSet[c.Point]()
	distance := sensor.distance

	dy := y - sensor.position.Y
	for dx := -distance + abs(dy); dx <= +distance-abs(dy); dx++ {
		occupied.Add(c.Point{X: sensor.position.X + dx, Y: y})
	}

	return occupied
}

func (sensor *Sensor) getBoundaries() (nw, ne, se, sw Line) {
	pos := sensor.position
	d := sensor.distance + 1

	left := c.Point{X: pos.X - d, Y: pos.Y}
	top := c.Point{X: pos.X, Y: pos.Y - d}
	right := c.Point{X: pos.X + d, Y: pos.Y}
	bottom := c.Point{X: pos.X, Y: pos.Y + d}

	nw[0] = left
	nw[1] = top

	ne[0] = top
	ne[1] = right

	se[0] = bottom
	se[1] = right

	sw[0] = left
	sw[1] = bottom

	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

type Line = [2]c.Point

func intersects(a, b Line) c.Point {
	aIntercept := a[0].Y - a[0].X
	bIntercept := b[0].Y + b[0].X

	x := (bIntercept - aIntercept) / 2
	y := (bIntercept + aIntercept) / 2

	return c.Point{X: x, Y: y}
}
