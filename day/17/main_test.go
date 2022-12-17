package main

import (
	. "common/cartesian"
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

	want := []Direction{
		Right, Right, Right, Left, Left, Right, Left, Right, Right, Left, Left, Left, Right, Right, Left, Right, Right, Right, Left, Left, Left, Right, Right, Right, Left, Left, Left, Right, Left, Left, Left, Right, Right, Left, Right, Right, Left, Left, Right, Right,
	}

	got := parseInput(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestTowerHeight(t *testing.T) {
	jets := []Direction{
		Right, Right, Right, Left, Left, Right, Left, Right, Right, Left, Left, Left, Right, Right, Left, Right, Right, Right, Left, Left, Left, Right, Right, Right, Left, Left, Left, Right, Left, Left, Left, Right, Right, Left, Right, Right, Left, Left, Right, Right,
	}

	tests := []struct {
		rocks int
		want  int
	}{
		{rocks: 1, want: 1},
		{rocks: 2, want: 4},
		{rocks: 3, want: 6},
		{rocks: 4, want: 7},
		{rocks: 5, want: 9},
		{rocks: 6, want: 10},
		{rocks: 7, want: 13},
		{rocks: 8, want: 15},
		{rocks: 9, want: 17},
		{rocks: 10, want: 17},
		{rocks: 2022, want: 3068},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("When [%d] rocks dropped, height is [%d]", tt.rocks, tt.want), func(t *testing.T) {
			got := getTowerHeight(jets, tt.rocks)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}
		})
	}

}
