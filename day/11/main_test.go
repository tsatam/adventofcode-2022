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
		want   []Monkey
	}{
		{
			rounds: 1,
			want: []Monkey{
				{
					items:     []int{20, 23, 27, 26},
					operation: MulOperation{19},
					testDiv:   23, throwTrue: 2, throwFalse: 3,

					inspected: 2,
				},
				{
					items:     []int{2080, 25, 167, 207, 401, 1046},
					operation: AddOperation{6},
					testDiv:   19, throwTrue: 2, throwFalse: 0,

					inspected: 4,
				},
				{
					items:     []int{},
					operation: PowOperation{},
					testDiv:   13, throwTrue: 1, throwFalse: 3,

					inspected: 3,
				},
				{
					items:     []int{},
					operation: AddOperation{3},
					testDiv:   17, throwTrue: 0, throwFalse: 1,

					inspected: 5,
				},
			},
		},
		{
			rounds: 2,
			want: []Monkey{
				{
					items:     []int{695, 10, 71, 135, 350},
					operation: MulOperation{19},
					testDiv:   23, throwTrue: 2, throwFalse: 3,

					inspected: 6,
				},
				{
					items:     []int{43, 49, 58, 55, 362},
					operation: AddOperation{6},
					testDiv:   19, throwTrue: 2, throwFalse: 0,

					inspected: 10,
				},
				{
					items:     []int{},
					operation: PowOperation{},
					testDiv:   13, throwTrue: 1, throwFalse: 3,

					inspected: 4,
				},
				{
					items:     []int{},
					operation: AddOperation{3},
					testDiv:   17, throwTrue: 0, throwFalse: 1,

					inspected: 10,
				},
			},
		},
		{
			rounds: 20,
			want: []Monkey{
				{
					items:     []int{10, 12, 14, 26, 34},
					operation: MulOperation{19},
					testDiv:   23, throwTrue: 2, throwFalse: 3,

					inspected: 101,
				},
				{
					items:     []int{245, 93, 53, 199, 115},
					operation: AddOperation{6},
					testDiv:   19, throwTrue: 2, throwFalse: 0,

					inspected: 95,
				},
				{
					items:     []int{},
					operation: PowOperation{},
					testDiv:   13, throwTrue: 1, throwFalse: 3,

					inspected: 7,
				},
				{
					items:     []int{},
					operation: AddOperation{3},
					testDiv:   17, throwTrue: 0, throwFalse: 1,

					inspected: 105,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rounds: [%d]", tt.rounds), func(t *testing.T) {
			got := processRounds(input(), tt.rounds)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got [\n%+v\n], want [\n%+v\n]", got, tt.want)
			}
		})
	}
}

func TestFindMonkeyBusiness(t *testing.T) {
	input := []Monkey{
		{
			items:     []int{10, 12, 14, 26, 34},
			operation: MulOperation{19},
			testDiv:   23, throwTrue: 2, throwFalse: 3,

			inspected: 101,
		},
		{
			items:     []int{245, 93, 53, 199, 115},
			operation: AddOperation{6},
			testDiv:   19, throwTrue: 2, throwFalse: 0,

			inspected: 95,
		},
		{
			items:     []int{},
			operation: PowOperation{},
			testDiv:   13, throwTrue: 1, throwFalse: 3,

			inspected: 7,
		},
		{
			items:     []int{},
			operation: AddOperation{3},
			testDiv:   17, throwTrue: 0, throwFalse: 1,

			inspected: 105,
		},
	}

	want := 10605
	got := findMonkeyBusiness(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
