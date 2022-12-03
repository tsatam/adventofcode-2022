package main

import "testing"

func TestProcessRucksack(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{name: "differnt items in both", line: "aA", want: 0},
		{name: "same item in both", line: "aa", want: 1},
		{name: "ignores not same items", line: "abccde", want: 3},
		{name: "example 1", line: "vJrwpWtwJgWrhcsFMMfFFhFp", want: 16},
		{name: "example 2", line: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", want: 38},
		{name: "example 3", line: "PmmdzqPrVvPwwTWBwg", want: 42},
		{name: "example 4", line: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", want: 22},
		{name: "example 5", line: "ttgJtRGJQctTZtZT", want: 20},
		{name: "example 6", line: "CrZsJsPPZsGzwwsLwLmpwMDw", want: 19},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			want := tt.want
			got := processRucksack(tt.line)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}

func TestProcessIndividually(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 157
	got := processIndividually(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestProcessInGroups(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 70
	got := processInGroups(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
