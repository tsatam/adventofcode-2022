package main

import (
	"testing"
)

func TestFindFirstNUnique(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantDistinct int
		want         int
	}{
		{
			name:         "unique from start",
			input:        "abcd",
			wantDistinct: 4,
			want:         4,
		},
		{
			name:         "ex0",
			input:        "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			wantDistinct: 4,
			want:         7,
		},
		{
			name:         "ex1",
			input:        "bvwbjplbgvbhsrlpgdmjqwftvncz",
			wantDistinct: 4,
			want:         5,
		},
		{
			name:         "ex2",
			input:        "nppdvjthqldpwncqszvftbrmjlhg",
			wantDistinct: 4,
			want:         6,
		},
		{
			name:         "ex3",
			input:        "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			wantDistinct: 4,
			want:         10,
		},
		{
			name:         "ex4",
			input:        "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			wantDistinct: 4,
			want:         11,
		},
		{
			name:         "ex0-pt2",
			input:        "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			wantDistinct: 14,
			want:         19,
		},
		{
			name:         "ex1-pt2",
			input:        "bvwbjplbgvbhsrlpgdmjqwftvncz",
			wantDistinct: 14,
			want:         23,
		},
		{
			name:         "ex2-pt2",
			input:        "nppdvjthqldpwncqszvftbrmjlhg",
			wantDistinct: 14,
			want:         23,
		},
		{
			name:         "ex3-pt2",
			input:        "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			wantDistinct: 14,
			want:         29,
		},
		{
			name:         "ex4-pt2",
			input:        "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			wantDistinct: 14,
			want:         26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			wantDistinct := tt.wantDistinct
			want := tt.want
			got := findFirstNUnique(input, wantDistinct)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}
