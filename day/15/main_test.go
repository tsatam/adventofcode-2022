package main

import (
	c "common/cartesian"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`
	want := []Sensor{
		{c.Point{X: 2, Y: 18}, c.Point{X: -2, Y: 15}, 7},
		{c.Point{X: 9, Y: 16}, c.Point{X: 10, Y: 16}, 1},
		{c.Point{X: 13, Y: 2}, c.Point{X: 15, Y: 3}, 3},
		{c.Point{X: 12, Y: 14}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 10, Y: 20}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 14, Y: 17}, c.Point{X: 10, Y: 16}, 5},
		{c.Point{X: 8, Y: 7}, c.Point{X: 2, Y: 10}, 9},
		{c.Point{X: 2, Y: 0}, c.Point{X: 2, Y: 10}, 10},
		{c.Point{X: 0, Y: 11}, c.Point{X: 2, Y: 10}, 3},
		{c.Point{X: 20, Y: 14}, c.Point{X: 25, Y: 17}, 8},
		{c.Point{X: 17, Y: 20}, c.Point{X: 21, Y: 22}, 6},
		{c.Point{X: 16, Y: 7}, c.Point{X: 15, Y: 3}, 5},
		{c.Point{X: 14, Y: 3}, c.Point{X: 15, Y: 3}, 1},
		{c.Point{X: 20, Y: 1}, c.Point{X: 15, Y: 3}, 7},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestGetCountOccupiedAt(t *testing.T) {
	sensors := []Sensor{
		{c.Point{X: 2, Y: 18}, c.Point{X: -2, Y: 15}, 7},
		{c.Point{X: 9, Y: 16}, c.Point{X: 10, Y: 16}, 1},
		{c.Point{X: 13, Y: 2}, c.Point{X: 15, Y: 3}, 3},
		{c.Point{X: 12, Y: 14}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 10, Y: 20}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 14, Y: 17}, c.Point{X: 10, Y: 16}, 5},
		{c.Point{X: 8, Y: 7}, c.Point{X: 2, Y: 10}, 9},
		{c.Point{X: 2, Y: 0}, c.Point{X: 2, Y: 10}, 10},
		{c.Point{X: 0, Y: 11}, c.Point{X: 2, Y: 10}, 3},
		{c.Point{X: 20, Y: 14}, c.Point{X: 25, Y: 17}, 8},
		{c.Point{X: 17, Y: 20}, c.Point{X: 21, Y: 22}, 6},
		{c.Point{X: 16, Y: 7}, c.Point{X: 15, Y: 3}, 5},
		{c.Point{X: 14, Y: 3}, c.Point{X: 15, Y: 3}, 1},
		{c.Point{X: 20, Y: 1}, c.Point{X: 15, Y: 3}, 7},
	}
	y := 10
	want := 26
	got := getCountOccupiedAt(sensors, y)

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestGetDistressBeaconTuningFrequency(t *testing.T) {
	sensors := []Sensor{
		{c.Point{X: 2, Y: 18}, c.Point{X: -2, Y: 15}, 7},
		{c.Point{X: 9, Y: 16}, c.Point{X: 10, Y: 16}, 1},
		{c.Point{X: 13, Y: 2}, c.Point{X: 15, Y: 3}, 3},
		{c.Point{X: 12, Y: 14}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 10, Y: 20}, c.Point{X: 10, Y: 16}, 4},
		{c.Point{X: 14, Y: 17}, c.Point{X: 10, Y: 16}, 5},
		{c.Point{X: 8, Y: 7}, c.Point{X: 2, Y: 10}, 9},
		{c.Point{X: 2, Y: 0}, c.Point{X: 2, Y: 10}, 10},
		{c.Point{X: 0, Y: 11}, c.Point{X: 2, Y: 10}, 3},
		{c.Point{X: 20, Y: 14}, c.Point{X: 25, Y: 17}, 8},
		{c.Point{X: 17, Y: 20}, c.Point{X: 21, Y: 22}, 6},
		{c.Point{X: 16, Y: 7}, c.Point{X: 15, Y: 3}, 5},
		{c.Point{X: 14, Y: 3}, c.Point{X: 15, Y: 3}, 1},
		{c.Point{X: 20, Y: 1}, c.Point{X: 15, Y: 3}, 7},
	}
	bounds := 20
	want := 56000011
	got := getDistressBeaconTuningFrequency(sensors, bounds)

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}

}

func TestIntersects(t *testing.T) {
	tests := []struct {
		a, b      Line
		wantOk    bool
		wantPoint c.Point
	}{
		{
			a:         Line{c.Point{X: 0, Y: 0}, c.Point{X: 4, Y: 4}},
			b:         Line{c.Point{X: 0, Y: 4}, c.Point{X: 4, Y: 0}},
			wantPoint: c.Point{X: 2, Y: 2},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotPoint := intersects(tt.a, tt.b)
			if gotPoint != tt.wantPoint {
				t.Errorf("got [%v], want [%v]", gotPoint, tt.wantPoint)
			}
		})
	}
}
