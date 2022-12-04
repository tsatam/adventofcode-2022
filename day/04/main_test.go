package main

import "testing"

func TestHasFullContain(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{line: "0-1,1-2", want: false},
		{line: "0-2,0-1", want: true},
		{line: "0-2,1-2", want: true},
		{line: "2-4,6-8", want: false},
		{line: "2-3,4-5", want: false},
		{line: "5-7,7-9", want: false},
		{line: "2-8,3-7", want: true},
		{line: "6-6,4-6", want: true},
		{line: "2-6,4-8", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			want := tt.want
			got := hasFullContain(tt.line)

			if got != want {
				t.Errorf("got [%v], want [%v]", got, want)
			}
		})
	}
}

func TestCountFullContains(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 2
	got := countFullContains(input)

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestHasOverlap(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{line: "0-1,1-2", want: true},
		{line: "0-2,0-1", want: true},
		{line: "0-2,1-2", want: true},
		{line: "2-4,6-8", want: false},
		{line: "2-3,4-5", want: false},
		{line: "5-7,7-9", want: true},
		{line: "2-8,3-7", want: true},
		{line: "6-6,4-6", want: true},
		{line: "2-6,4-8", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			want := tt.want
			got := hasOverlap(tt.line)

			if got != want {
				t.Errorf("got [%v], want [%v]", got, want)
			}
		})
	}
}

func TestCountOverlaps(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 4
	got := countOverlaps(input)

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}
