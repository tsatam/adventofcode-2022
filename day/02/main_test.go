package main

import (
	"testing"
)

func TestProcessLinePt1(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "A X", want: 4},
		{line: "A Y", want: 8},
		{line: "A Z", want: 3},
		{line: "B X", want: 1},
		{line: "B Y", want: 5},
		{line: "B Z", want: 9},
		{line: "C X", want: 7},
		{line: "C Y", want: 2},
		{line: "C Z", want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			want := tt.want
			got := processLinePt1(tt.line)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}

func TestProcessLinePt2(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "A X", want: 3},
		{line: "A Y", want: 4},
		{line: "A Z", want: 8},
		{line: "B X", want: 1},
		{line: "B Y", want: 5},
		{line: "B Z", want: 9},
		{line: "C X", want: 2},
		{line: "C Y", want: 6},
		{line: "C Z", want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			want := tt.want
			got := processLinePt2(tt.line)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}

func TestProcessGamePt1(t *testing.T) {
	input := `A Y
B X
C Z`

	want := 15
	got, _ := processGame(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestProcessGamePt2(t *testing.T) {
	input := `A Y
B X
C Z`

	want := 12
	_, got := processGame(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
