package main

import (
	c "common/cartesian"
	mpq "common/min_priority_queue"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	start, dest, heightmap := parseInput(input)

	shortestPathFromStart, shortestPaths := findShortestPath(start, dest, heightmap)
	shortestPathFromAnyLowestPoint := findShortestPathFromAnyLowestPoint(heightmap, shortestPaths)

	fmt.Printf("pt1: [%d]\n", shortestPathFromStart)
	fmt.Printf("pt2: [%d]\n", shortestPathFromAnyLowestPoint)
}

func parseInput(input string) (start c.Point, dest c.Point, heightmap [][]int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	heightmap = make([][]int, len(lines))

	for y, line := range lines {
		heightmap[y] = make([]int, len(line))
		for x, char := range line {
			if char == 'S' {
				start = c.Point{X: x, Y: y}
				char = 'a'
			}
			if char == 'E' {
				dest = c.Point{X: x, Y: y}
				char = 'z'
			}
			heightmap[y][x] = int(char - 'a')
		}
	}

	return
}

func findShortestPath(start, dest c.Point, heightmap [][]int) (uint, [][]uint) {
	boundsY, boundsX := len(heightmap), len(heightmap[0])

	queue := mpq.New[c.Point](0, boundsX*boundsY)
	shortestDistances := make([][]uint, boundsY)

	for y := range heightmap {
		shortestDistances[y] = make([]uint, boundsX)
		for x := range heightmap[y] {
			p := c.Point{X: x, Y: y}
			if dest.X == x && dest.Y == y {
				shortestDistances[y][x] = 0
			} else {
				shortestDistances[y][x] = math.MaxInt
			}
			queue.AddAtPriority(p, int(shortestDistances[y][x]))
		}
	}

	validMove := func(p c.Point, d c.Direction) (c.Point, bool) {
		res := p.Move(d)

		if res.X < 0 || res.Y < 0 || res.X >= boundsX || res.Y >= boundsY {
			return c.Point{X: -1, Y: -1}, false
		}
		elevationDifference := heightmap[p.Y][p.X] - heightmap[res.Y][res.X]
		if elevationDifference > 1 {
			return c.Point{X: -1, Y: -1}, false
		}

		return res, true
	}

	getNeighbors := func(p c.Point) []c.Point {
		res := make([]c.Point, 0, 4)

		for _, d := range []c.Direction{c.Up, c.Down, c.Left, c.Right} {
			neighbor, ok := validMove(p, d)
			if ok {
				res = append(res, neighbor)
			}
		}
		return res
	}

	for !queue.Empty() {
		next := queue.PopMin()

		for _, neighbor := range getNeighbors(next) {
			currentDistance := shortestDistances[next.Y][next.X] + 1

			if currentDistance < shortestDistances[neighbor.Y][neighbor.X] {
				shortestDistances[neighbor.Y][neighbor.X] = currentDistance
				queue.SetPriority(neighbor, int(currentDistance))
			}
		}
	}

	return shortestDistances[start.Y][start.X], shortestDistances
}

func findShortestPathFromAnyLowestPoint(heightmap [][]int, shortestPath [][]uint) uint {
	minShortestPath := uint(math.MaxInt - 1)
	for y := range heightmap {
		for x := range heightmap[y] {
			if heightmap[y][x] == 0 && shortestPath[y][x] < minShortestPath {
				minShortestPath = shortestPath[y][x]
			}
		}
	}
	return minShortestPath
}
