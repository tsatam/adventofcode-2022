package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name: "single tree",
			input: `5
`,
			want: [][]int{{5}},
		},
		{
			name: "example",
			input: `30373
25512
65332
33549
35390
`,
			want: [][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := parseInput(input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got [%v], want [%v]", got, want)
			}
		})
	}
}

func TestFindVisibleTrees(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name:  "single tree",
			input: [][]int{{5}},
			want:  [][]int{{1}},
		},
		{
			name: "example",
			input: [][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			},
			want: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 0, 1, 1},
				{1, 0, 1, 0, 1},
				{1, 1, 1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := findVisibleTrees(input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got [%v], want [%v]", got, want)
			}
		})
	}
}

func TestSumOfVisibleTrees(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name:  "single tree",
			input: [][]int{{1}},
			want:  1,
		},
		{
			name: "example",
			input: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 0, 1, 1},
				{1, 0, 1, 0, 1},
				{1, 1, 1, 1, 1},
			},
			want: 21,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := sumOfVisibleTrees(input)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}
