package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	dirSizes := getDirSizes(input)

	fmt.Printf("pt1: %d\n", getSumOfAllDirectoriesLessThan(dirSizes, 100000))
	fmt.Printf("pt2: %d\n", getSizeOfSmallestDirToDelete(dirSizes, 70000000, 30000000))
}

func getDirSizes(input string) map[string]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	dirSizes := map[string]int{}
	dirStack := make([]string, 0)

	for _, line := range lines {
		if line == "$ ls" || strings.HasPrefix(line, "dir ") {
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			dirName := line[5:]
			if dirName == ".." {
				dirStack = dirStack[:len(dirStack)-1]
			} else {
				if dirName != "/" {
					dirName = dirName + "/"
				}
				dirStack = append(dirStack, dirName)
				dirSizes[buildDirPath(dirStack)] = 0
			}
		} else {
			splitLine := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(splitLine[0])

			for i := range dirStack {
				dirSizes[buildDirPath(dirStack[:i+1])] += fileSize
			}
		}
	}

	return dirSizes
}

func getSumOfAllDirectoriesLessThan(dirSizes map[string]int, bound int) int {
	sum := 0
	for _, size := range dirSizes {
		if size <= bound {
			sum += size
		}
	}
	return sum
}

func getSizeOfSmallestDirToDelete(dirSizes map[string]int, total int, needFree int) int {
	used := dirSizes["/"]
	haveFree := total - used
	shouldFree := needFree - haveFree

	smallestFree := used
	for _, size := range dirSizes {
		if size >= shouldFree && size < smallestFree {
			smallestFree = size
		}
	}
	return smallestFree
}

func buildDirPath(dirStack []string) string {
	return strings.Join(dirStack, "")
}
