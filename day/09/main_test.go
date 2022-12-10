package main

import (
	cartesian "common/cartesian"
	"fmt"
	"reflect"
	"testing"
)

func TestParseInstructions(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	want := []Motion{
		{cartesian.Right, 4},
		{cartesian.Up, 4},
		{cartesian.Left, 3},
		{cartesian.Down, 1},
		{cartesian.Right, 4},
		{cartesian.Down, 1},
		{cartesian.Left, 5},
		{cartesian.Right, 2},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestReconcileTail(t *testing.T) {
	tests := []struct {
		head cartesian.Point
		tail cartesian.Point
		want cartesian.Point
	}{
		{cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},

		{cartesian.Point{X: -1, Y: 0}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: -1, Y: -1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: 0, Y: -1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: 1, Y: -1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: 1, Y: 0}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: 1, Y: 1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: 0, Y: 1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},
		{cartesian.Point{X: -1, Y: 1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 0}},

		{cartesian.Point{X: -2, Y: 0}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: -1, Y: 0}},
		{cartesian.Point{X: -2, Y: -1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: -1, Y: -1}},
		{cartesian.Point{X: -2, Y: 1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: -1, Y: 1}},

		{cartesian.Point{X: 0, Y: -2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: -1}},
		{cartesian.Point{X: -1, Y: -2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: -1, Y: -1}},
		{cartesian.Point{X: 1, Y: -2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 1, Y: -1}},

		{cartesian.Point{X: 2, Y: 0}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 1, Y: 0}},
		{cartesian.Point{X: 2, Y: -1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 1, Y: -1}},
		{cartesian.Point{X: 2, Y: 1}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 1, Y: 1}},

		{cartesian.Point{X: 0, Y: 2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 0, Y: 1}},
		{cartesian.Point{X: -1, Y: 2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: -1, Y: 1}},
		{cartesian.Point{X: 1, Y: 2}, cartesian.Point{X: 0, Y: 0}, cartesian.Point{X: 1, Y: 1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Head: %v, Tail: %v -> %v", tt.head, tt.tail, tt.want), func(t *testing.T) {
			got := reconcileTail(tt.head, tt.tail)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got [%v], want [%v]", got, tt.want)
			}
		})
	}
}

func TestPositionsVisited(t *testing.T) {
	motions := []Motion{
		{cartesian.Right, 4},
		{cartesian.Up, 4},
		{cartesian.Left, 3},
		{cartesian.Down, 1},
		{cartesian.Right, 4},
		{cartesian.Down, 1},
		{cartesian.Left, 5},
		{cartesian.Right, 2},
	}
	want := 13
	got := positionsVisited(motions, 2)
	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestPositionsVisitedPt2(t *testing.T) {
	motions := []Motion{
		{cartesian.Right, 4},
		{cartesian.Up, 4},
		{cartesian.Left, 3},
		{cartesian.Down, 1},
		{cartesian.Right, 4},
		{cartesian.Down, 1},
		{cartesian.Left, 5},
		{cartesian.Right, 2},
	}
	want := 1
	got := positionsVisited(motions, 10)
	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
