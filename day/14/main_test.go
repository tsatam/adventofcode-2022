package main

import (
	c "common/cartesian"
	s "common/set"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`
	want := []Line{
		{{X: 498, Y: 4}, {X: 498, Y: 6}, {X: 496, Y: 6}},
		{{X: 503, Y: 4}, {X: 502, Y: 4}, {X: 502, Y: 9}, {X: 494, Y: 9}},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestGetPointsForLines(t *testing.T) {
	input := []Line{
		{{X: 498, Y: 4}, {X: 498, Y: 6}, {X: 496, Y: 6}},
		{{X: 503, Y: 4}, {X: 502, Y: 4}, {X: 502, Y: 9}, {X: 494, Y: 9}},
	}

	want := s.NewSet[c.Point]()
	want.AddAll([]c.Point{
		{X: 498, Y: 4},
		{X: 498, Y: 5},
		{X: 498, Y: 6},
		{X: 497, Y: 6},
		{X: 496, Y: 6},

		{X: 503, Y: 4},
		{X: 502, Y: 4},
		{X: 502, Y: 5}, {X: 502, Y: 6}, {X: 502, Y: 7}, {X: 502, Y: 8},
		{X: 502, Y: 9},
		{X: 501, Y: 9}, {X: 500, Y: 9}, {X: 499, Y: 9}, {X: 498, Y: 9}, {X: 497, Y: 9}, {X: 496, Y: 9}, {X: 495, Y: 9},
		{X: 494, Y: 9},
	}...)

	got := getPointsForLines(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestProcessSandNoFloor(t *testing.T) {
	input := s.NewSet[c.Point]()
	input.AddAll([]c.Point{
		{X: 498, Y: 4},
		{X: 498, Y: 5},
		{X: 498, Y: 6},
		{X: 497, Y: 6},
		{X: 496, Y: 6},

		{X: 503, Y: 4},
		{X: 502, Y: 4},
		{X: 502, Y: 5}, {X: 502, Y: 6}, {X: 502, Y: 7}, {X: 502, Y: 8},
		{X: 502, Y: 9},
		{X: 501, Y: 9}, {X: 500, Y: 9}, {X: 499, Y: 9}, {X: 498, Y: 9}, {X: 497, Y: 9}, {X: 496, Y: 9}, {X: 495, Y: 9},
		{X: 494, Y: 9},
	}...)

	want := 24
	got := processSand(input, false)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestProcessSandFloor(t *testing.T) {
	input := s.NewSet[c.Point]()
	input.AddAll([]c.Point{
		{X: 498, Y: 4},
		{X: 498, Y: 5},
		{X: 498, Y: 6},
		{X: 497, Y: 6},
		{X: 496, Y: 6},

		{X: 503, Y: 4},
		{X: 502, Y: 4},
		{X: 502, Y: 5}, {X: 502, Y: 6}, {X: 502, Y: 7}, {X: 502, Y: 8},
		{X: 502, Y: 9},
		{X: 501, Y: 9}, {X: 500, Y: 9}, {X: 499, Y: 9}, {X: 498, Y: 9}, {X: 497, Y: 9}, {X: 496, Y: 9}, {X: 495, Y: 9},
		{X: 494, Y: 9},
	}...)

	want := 93
	got := processSand(input, true)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
