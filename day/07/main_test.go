package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Path
	}{
		{
			name: "root with single file",
			input: `$ cd /
$ ls
10 a.txt
`,
			want: Dir{
				name:  "/",
				paths: []Path{File{"a.txt", 10}},
			},
		},
		{
			name: "example",
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`,
			want: Dir{"/", []Path{
				Dir{"a", []Path{
					Dir{"e", []Path{File{"i", 584}}},
					File{"f", 29116},
					File{"g", 2557},
					File{"h.lst", 62596},
				}},
				File{"b.txt", 14848514},
				File{"c.dat", 8504156},
				Dir{"d", []Path{
					File{"j", 4060174},
					File{"d.log", 8033020},
					File{"d.ext", 5626152},
					File{"k", 7214296},
				}},
			}},
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

func TestSize(t *testing.T) {
	tests := []struct {
		name  string
		input Path
		want  int
	}{
		{
			name:  "simple file",
			input: File{"a", 10},
			want:  10,
		},
		{
			name:  "simple dir",
			input: Dir{"/", []Path{File{"a", 10}}},
			want:  10,
		},
		{
			name:  "multiple files in dir",
			input: Dir{"/", []Path{File{"a", 0}, File{"b", 20}}},
			want:  30,
		},
		{
			name:  "nested dir",
			input: Dir{"/", []Path{Dir{"x", []Path{File{"a", 10}, File{"b", 20}}}, File{"c", 30}}},
			want:  60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := input.Size()

			if got != want {
				t.Errorf("got [%d], want [%d]", got, want)
			}
		})
	}
}

func TestSumOfAllDirectoriesLE100000(t *testing.T) {
	tests := []struct {
		name  string
		input Dir
		want  int
	}{
		{
			name:  "simple dir",
			input: Dir{"/", []Path{File{"a", 10}}},
			want:  10,
		},
		{
			name:  "nested dir",
			input: Dir{"/", []Path{Dir{"x", []Path{File{"a", 10}, File{"b", 20}}}, File{"c", 30}}},
			want:  90,
		},
		{
			name: "example",
			input: Dir{"/", []Path{
				Dir{"a", []Path{
					Dir{"e", []Path{File{"i", 584}}},
					File{"f", 29116},
					File{"g", 2557},
					File{"h.lst", 62596},
				}},
				File{"b.txt", 14848514},
				File{"c.dat", 8504156},
				Dir{"d", []Path{
					File{"j", 4060174},
					File{"d.log", 8033020},
					File{"d.ext", 5626152},
					File{"k", 7214296},
				}},
			}},
			want: 95437,
		},
	}

	for _, tt := range tests {
		input := tt.input
		want := tt.want
		got := input.SumOfAllDirectoriesLE100000()

		if got != want {
			t.Errorf("got [%d], want [%d]", got, want)
		}
	}
}
