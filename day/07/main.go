package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	root := parseInput(input)
	bound := 30000000 - (70000000 - root.Size())

	fmt.Printf("pt1: %d\n", root.SumOfAllDirectoriesLE100000())
	fmt.Printf("pt2: %d\n", root.FindSmallestDirGE(bound))
}

func parseInput(input string) Dir {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	root := Dir{
		name:  "/",
		paths: []Path{},
	}

	i := 0

	root = parseDir(root, lines, &i)

	return root
}

func parseDir(dir Dir, lines []string, i *int) Dir {
	pathsToAdd := map[string]Path{}

	if lines[*i] != fmt.Sprintf("$ cd %s", dir.Name()) {
		log.Fatalf("wrong line, expected cd. i [%d], dir [%v], line [%s]", *i, dir, lines[*i])
	}
	*i++
	if lines[*i] != "$ ls" {
		log.Fatal("wrong line, expected ls")
	}
	*i++

	for ; *i < len(lines) && !strings.HasPrefix(lines[*i], "$ cd"); *i++ {
		path := getPathFromLsLine(lines[*i])
		pathsToAdd[path.Name()] = path
	}

	for ; *i < len(lines); *i++ {
		nextDir := getDirNameFromCd(lines[*i])

		if nextDir == ".." {
			break
		} else {
			pathsToAdd[nextDir] = parseDir(pathsToAdd[nextDir].(Dir), lines, i)
		}
	}

	for _, v := range pathsToAdd {
		dir.paths = append(dir.paths, v)
	}
	return dir
}

func getDirNameFromCd(line string) string {
	var dirName string
	if _, err := fmt.Sscanf(line, "$ cd %s", &dirName); err != nil {
		log.Fatal(err)
	}
	return dirName
}

func getPathFromLsLine(line string) Path {
	if strings.HasPrefix(line, "$") {
		log.Fatal("parsing wrong ls: " + line)
	}

	if strings.HasPrefix(line, "dir ") {
		var name string
		if _, err := fmt.Sscanf(line, "dir %s", &name); err != nil {
			log.Fatal(err)
		}
		return Dir{
			name:  name,
			paths: []Path{},
		}
	} else {
		var name string
		var size int
		if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err != nil {
			log.Fatal(err)
		}
		return File{
			name: name,
			size: size,
		}
	}
}

type Path interface {
	Name() string
	Size() int
}

type File struct {
	name string
	size int
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}

type Dir struct {
	name  string
	paths []Path
}

func (d Dir) Name() string {
	return d.name
}

func (d Dir) Size() int {
	sum := 0
	for _, path := range d.paths {
		sum += path.Size()
	}
	return sum
}

func (d Dir) SumOfAllDirectoriesLE100000() int {
	sum := 0
	for _, p := range d.paths {
		if sp, ok := p.(Dir); ok {
			sum += sp.SumOfAllDirectoriesLE100000()
		}
	}
	if d.Size() <= 100000 {
		sum += d.Size()
	}
	return sum
}

func (d Dir) FindSmallestDirGE(bound int) int {
	smallestSize := math.MaxInt
	if d.Size() >= bound {
		smallestSize = d.Size()
	}
	for _, p := range d.paths {
		if sp, ok := p.(Dir); ok {
			size := sp.FindSmallestDirGE(bound)

			if size >= bound && size < smallestSize {
				smallestSize = size
			}
		}
	}

	return smallestSize
}
