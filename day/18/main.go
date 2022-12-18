package main

import (
	s "common/set"
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
	cubes := parseInput(input)
	surfaceArea := calcSurfaceArea(cubes)

	fmt.Printf("pt1: [%d]\n", surfaceArea)
}

func parseInput(input string) []Point3D {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	points := make([]Point3D, len(lines))

	for i, line := range lines {
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points[i] = Point3D{x, y, z}
	}
	return points
}

func calcSurfaceArea(cubes []Point3D) int {
	set := s.NewSet(cubes...)

	surfaceArea := 0

	for _, cube := range cubes {
		for _, adjacent := range cube.getAdjacent() {
			if !set.Contains(adjacent) {
				surfaceArea++
			}
		}
	}

	return surfaceArea
}

type Point3D struct {
	X, Y, Z int
}

func (p *Point3D) getAdjacent() []Point3D {
	return []Point3D{
		{p.X, p.Y, p.Z - 1},
		{p.X, p.Y, p.Z + 1},
		{p.X, p.Y - 1, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X - 1, p.Y, p.Z},
		{p.X + 1, p.Y, p.Z},
	}
}
