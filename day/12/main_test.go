package main

import (
	c "common/cartesian"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`
	wantStart, wantDest, wantMap := c.Point{X: 0, Y: 0}, c.Point{X: 5, Y: 2}, [][]int{
		{0, 0, 1, 16, 15, 14, 13, 12},
		{0, 1, 2, 17, 24, 23, 23, 11},
		{0, 2, 2, 18, 25, 25, 23, 10},
		{0, 2, 2, 19, 20, 21, 22, 9},
		{0, 1, 3, 4, 5, 6, 7, 8},
	}

	gotStart, gotDest, gotMap := parseInput(input)

	if !reflect.DeepEqual(gotStart, wantStart) {
		t.Errorf("got [%v], want [%v]", gotStart, wantStart)
	}

	if !reflect.DeepEqual(gotDest, wantDest) {
		t.Errorf("got [%v], want [%v]", gotDest, wantDest)
	}
	if !reflect.DeepEqual(gotMap, wantMap) {
		t.Errorf("got [%v], want [%v]", gotMap, wantMap)
	}
}

func TestFindShortestPath(t *testing.T) {
	tests := []struct {
		name      string
		start     c.Point
		dest      c.Point
		heightmap [][]int
		want      int
	}{
		{
			name:      "1 row",
			start:     c.Point{X: 0, Y: 0},
			dest:      c.Point{X: 25, Y: 0},
			heightmap: [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
			want:      25,
		},
		{
			name:  "example",
			start: c.Point{X: 0, Y: 0},
			dest:  c.Point{X: 5, Y: 2},
			heightmap: [][]int{
				{0, 0, 1, 16, 15, 14, 13, 12},
				{0, 1, 2, 17, 24, 23, 23, 11},
				{0, 2, 2, 18, 25, 25, 23, 10},
				{0, 2, 2, 19, 20, 21, 22, 9},
				{0, 1, 3, 4, 5, 6, 7, 8},
			},
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findShortestPath(tt.start, tt.dest, tt.heightmap)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}

		})
	}
}

func TestFindShortestPathFromAnyLowestPoint(t *testing.T) {
	tests := []struct {
		name      string
		dest      c.Point
		heightmap [][]int
		want      int
	}{
		{
			name: "example",
			dest: c.Point{X: 5, Y: 2},
			heightmap: [][]int{
				{0, 0, 1, 16, 15, 14, 13, 12},
				{0, 1, 2, 17, 24, 23, 23, 11},
				{0, 2, 2, 18, 25, 25, 23, 10},
				{0, 2, 2, 19, 20, 21, 22, 9},
				{0, 1, 3, 4, 5, 6, 7, 8},
			},
			want: 29,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findShortestPathFromAnyLowestPoint(tt.dest, tt.heightmap)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}

		})
	}
}
