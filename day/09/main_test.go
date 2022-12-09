package main

import (
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
		{Right, 4},
		{Up, 4},
		{Left, 3},
		{Down, 1},
		{Right, 4},
		{Down, 1},
		{Left, 5},
		{Right, 2},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestReconcileTail(t *testing.T) {
	tests := []struct {
		head Point
		tail Point
		want Point
	}{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},

		{Point{-1, 0}, Point{0, 0}, Point{0, 0}},
		{Point{-1, -1}, Point{0, 0}, Point{0, 0}},
		{Point{0, -1}, Point{0, 0}, Point{0, 0}},
		{Point{1, -1}, Point{0, 0}, Point{0, 0}},
		{Point{1, 0}, Point{0, 0}, Point{0, 0}},
		{Point{1, 1}, Point{0, 0}, Point{0, 0}},
		{Point{0, 1}, Point{0, 0}, Point{0, 0}},
		{Point{-1, 1}, Point{0, 0}, Point{0, 0}},

		{Point{-2, 0}, Point{0, 0}, Point{-1, 0}},
		{Point{-2, -1}, Point{0, 0}, Point{-1, -1}},
		{Point{-2, 1}, Point{0, 0}, Point{-1, 1}},

		{Point{0, -2}, Point{0, 0}, Point{0, -1}},
		{Point{-1, -2}, Point{0, 0}, Point{-1, -1}},
		{Point{1, -2}, Point{0, 0}, Point{1, -1}},

		{Point{2, 0}, Point{0, 0}, Point{1, 0}},
		{Point{2, -1}, Point{0, 0}, Point{1, -1}},
		{Point{2, 1}, Point{0, 0}, Point{1, 1}},

		{Point{0, 2}, Point{0, 0}, Point{0, 1}},
		{Point{-1, 2}, Point{0, 0}, Point{-1, 1}},
		{Point{1, 2}, Point{0, 0}, Point{1, 1}},
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
		{Right, 4},
		{Up, 4},
		{Left, 3},
		{Down, 1},
		{Right, 4},
		{Down, 1},
		{Left, 5},
		{Right, 2},
	}
	want := 13
	got := positionsVisited(motions, 2)
	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestPositionsVisitedPt2(t *testing.T) {
	motions := []Motion{
		{Right, 4},
		{Up, 4},
		{Left, 3},
		{Down, 1},
		{Right, 4},
		{Down, 1},
		{Left, 5},
		{Right, 2},
	}
	want := 1
	got := positionsVisited(motions, 10)
	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
