package main

import (
	"reflect"
	"testing"
)

func TestGetDirSizes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name: "root with single file",
			input: `$ cd /
$ ls
10 a.txt
`,
			want: map[string]int{"/": 10},
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
			want: map[string]int{
				"/":     48381165,
				"/a/":   94853,
				"/a/e/": 584,
				"/d/":   24933642,
			},
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			input := tt.input
			want := tt.want
			got := getDirSizes(input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got [%v], want [%v]", got, want)
			}
		})
	}
}

func TestGetSumofAllDirectoriesLessThan(t *testing.T) {
	tests := []struct {
		name     string
		dirSizes map[string]int
		bound    int
		want     int
	}{
		{
			name:     "root with single file",
			dirSizes: map[string]int{"/": 10},
			bound:    100,
			want:     10,
		},
		{
			name: "example",
			dirSizes: map[string]int{
				"/":     48381165,
				"/a/":   94853,
				"/a/e/": 584,
				"/d/":   24933642,
			},
			bound: 100000,
			want:  95437,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			got := getSumOfAllDirectoriesLessThan(tt.dirSizes, tt.bound)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}
		})
	}
}

func TestGetSizeOfSmallestDirToDelete(t *testing.T) {
	tests := []struct {
		name     string
		dirSizes map[string]int
		total    int
		needFree int
		want     int
	}{
		{
			name:     "root with single file",
			dirSizes: map[string]int{"/": 10},
			total:    10,
			needFree: 10,
			want:     10,
		},
		{
			name: "example",
			dirSizes: map[string]int{
				"/":     48381165,
				"/a/":   94853,
				"/a/e/": 584,
				"/d/":   24933642,
			},
			total:    70000000,
			needFree: 30000000,
			want:     24933642,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			got := getSizeOfSmallestDirToDelete(tt.dirSizes, tt.total, tt.needFree)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}
		})
	}
}

// old impl: 528879 ns/op	  152719 B/op	    5699 allocs/op
// new impl: 138912 ns/op	   90640 B/op	    2067 allocs/op
func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
