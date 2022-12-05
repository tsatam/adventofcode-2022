package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	input := `[A]
1 

move 1 from 1 to 1`
	want := Ship{
		crates: [][]byte{{'A'}},
		instructions: []Instruction{
			{1, 1, 1},
		},
	}
	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestProcessInstructions9000(t *testing.T) {
	input := Ship{
		crates: [][]byte{
			{'Z', 'N'},
			{'M', 'C', 'D'},
			{'P'},
		},
		instructions: []Instruction{
			{1, 2, 1},
			{3, 1, 3},
			{2, 2, 1},
			{1, 1, 2},
		},
	}
	want := "CMZ"
	input.processInstructions9000()
	got := input.outputHighestCrates()

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestProcessInstructions9001(t *testing.T) {
	input := Ship{
		crates: [][]byte{
			{'Z', 'N'},
			{'M', 'C', 'D'},
			{'P'},
		},
		instructions: []Instruction{
			{1, 2, 1},
			{3, 1, 3},
			{2, 2, 1},
			{1, 1, 2},
		},
	}
	want := "MCD"
	input.processInstructions9001()
	got := input.outputHighestCrates()

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}
