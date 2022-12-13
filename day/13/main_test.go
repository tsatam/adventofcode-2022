package main

import (
	"fmt"
	"testing"
)

func TestCompareSignals(t *testing.T) {
	tests := []struct {
		left, right string
		want        bool
	}{
		{"1", "1", true},
		{"1", "1", true},
		{"2", "1", false},

		{"1", "[1]", true},
		{"[1]", "1", true},
		{"2", "[1]", false},

		{"[1,1,3,1,1]", "[1,1,5,1,1]", true},
		{"[[1],[2,3,4]]", "[[1],4]", true},
		{"[9]", "[[8,7,6]]", false},
		{"[[4,4],4,4]", "[[4,4],4,4,4]", true},
		{"[7,7,7,7]", "[7,7,7]", false},
		{"[]", "[3]", true},
		{"[[[]]]", "[[]]", false},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s<>%s=>[%v]", tt.left, tt.right, tt.want), func(t *testing.T) {
			got := compareSignals(tt.left, tt.right)

			if got != tt.want {
				t.Errorf("got [%v], want [%v]", got, tt.want)
			}
		})
	}

}

func TestSolve(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

	want := 13
	got := solve(input)

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}
