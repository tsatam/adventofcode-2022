package main

import (
	"testing"
)

func TestReadInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "1 elf with 1 food",
			input: `100
`,
			want: 100,
		},
		{
			name: "3 elves with 1 food, picks highest",
			input: `100

300

200`,
			want: 300,
		},
		{
			name: "1 elf with 3 food, sums calories",
			input: `100
200
300
`,
			want: 600,
		},
		{
			name: "example",
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000
`,
			want: 24000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := readInput(input)

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}
