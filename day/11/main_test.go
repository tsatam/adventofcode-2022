package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	input := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

	want := []Monkey{
		{
			items:     []int{79, 98},
			operation: MulOperation{19},
			testDiv:   23, throwTrue: 2, throwFalse: 3,

			inspected: 0,
		},
		{
			items:     []int{54, 65, 75, 74},
			operation: AddOperation{6},
			testDiv:   19, throwTrue: 2, throwFalse: 0,

			inspected: 0,
		},
		{
			items:     []int{79, 60, 97},
			operation: PowOperation{},
			testDiv:   13, throwTrue: 1, throwFalse: 3,

			inspected: 0,
		},
		{
			items:     []int{74},
			operation: AddOperation{3},
			testDiv:   17, throwTrue: 0, throwFalse: 1,

			inspected: 0,
		},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestProcessRounds(t *testing.T) {

	input := func() []Monkey {
		return []Monkey{
			{
				items:     []int{79, 98},
				operation: MulOperation{19},
				testDiv:   23, throwTrue: 2, throwFalse: 3,

				inspected: 0,
			},
			{
				items:     []int{54, 65, 75, 74},
				operation: AddOperation{6},
				testDiv:   19, throwTrue: 2, throwFalse: 0,

				inspected: 0,
			},
			{
				items:     []int{79, 60, 97},
				operation: PowOperation{},
				testDiv:   13, throwTrue: 1, throwFalse: 3,

				inspected: 0,
			},
			{
				items:     []int{74},
				operation: AddOperation{3},
				testDiv:   17, throwTrue: 0, throwFalse: 1,

				inspected: 0,
			},
		}
	}

	tests := []struct {
		rounds int
		want   []int
		worry  bool
	}{
		{
			rounds: 1,
			want:   []int{2, 4, 3, 5},
		},
		{
			rounds: 2,
			want:   []int{6, 10, 4, 10},
		},
		{
			rounds: 20,
			want:   []int{101, 95, 7, 105},
		},
		{
			rounds: 1,
			worry:  true,
			want:   []int{2, 4, 3, 6},
		},
		{
			rounds: 20,
			worry:  true,
			want:   []int{99, 97, 8, 103},
		},
		{
			rounds: 1000,
			worry:  true,
			want:   []int{5204, 4792, 199, 5192},
		},
		{
			rounds: 10000,
			worry:  true,
			want:   []int{52166, 47830, 1938, 52013},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rounds: [%d], Worry [%v]", tt.rounds, tt.worry), func(t *testing.T) {
			got := processRounds(input(), tt.rounds, tt.worry)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got [\n%+v\n], want [\n%+v\n]", got, tt.want)
			}
		})
	}
}

func TestFindMonkeyBusiness(t *testing.T) {
	input := []int{101, 95, 7, 105}

	want := int(10605)
	got := findMonkeyBusiness(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
