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
}

type Sensor struct {
	position      c.Point
	closestBeacon c.Point
}

func parseInput(input string) []Sensor {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sensors := make([]Sensor, len(lines))

	for i, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int

		if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY); err != nil {
			log.Fatal(err)
		}

		sensors[i] = Sensor{
			position:      c.Point{X: sensorX, Y: sensorY},
			closestBeacon: c.Point{X: beaconX, Y: beaconY},
		}
	}

	return sensors
}

func getOccupiedPositions(sensors []Sensor, y int) s.Set[c.Point] {
	occupied := s.NewSet[c.Point]()

	for _, sensor := range sensors {
		occupied.Merge(sensor.getOccupied(y))
	}
	for _, sensor := range sensors {
		occupied.Remove(sensor.closestBeacon)
	}
	return occupied
}

func getCountOccupiedAt(sensors []Sensor, y int) int {
	occupiedSet := getOccupiedPositions(sensors, y)
	occupied := occupiedSet.Slice()
	// sort.Slice(occupied, func(i, j int) bool { return occupied[i].X < occupied[j].X }) // debug helper

	sum := 0

	for _, point := range occupied {
		if point.Y == y {
			sum++
		}
	}
	return sum
}

func (sensor *Sensor) getBeaconDistance() int {
	return abs(sensor.position.X-sensor.closestBeacon.X) +
		abs(sensor.position.Y-sensor.closestBeacon.Y)
}

func (sensor *Sensor) getOccupied(y int) s.Set[c.Point] {
	occupied := s.NewSet[c.Point]()
	distance := sensor.getBeaconDistance()

	if abs(sensor.position.Y-y) > distance {
		return occupied
	}

	for x := -distance + abs(y-sensor.position.Y); x <= distance-abs(y-sensor.position.Y); x++ {
		occupied.Add(c.Point{X: sensor.position.X + x, Y: y})
	}

	return occupied
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
