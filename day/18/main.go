package main

import (
	minpq "common/min_priority_queue"
	s "common/set"
	_ "embed"
	"fmt"
	"math"
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

	surfaceAreaNoInterior := calcSurfaceAreaNoInterior(cubes)

	fmt.Printf("pt2: [%d]\n", surfaceAreaNoInterior)
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

func calcSurfaceAreaNoInterior(cubes []Point3D) int {
	// allCubes := addInteriorCubes(cubes)
	// return calcSurfaceArea(allCubes)

	boundaryCubes := s.NewSet(cubes...)

	min, max := getMinMax(cubes)

	surfaceArea := 0

	for _, cube := range cubes {
		for _, adjacent := range cube.getAdjacent() {
			if !boundaryCubes.Contains(adjacent) {
				isInterior, checkedAdjacent := isCubeInterior(adjacent, boundaryCubes, min, max)
				if !isInterior {
					surfaceArea++
				} else {
					boundaryCubes.Union(checkedAdjacent)
				}
			}
		}
	}

	return surfaceArea
}

func getPriority(cube, min, max Point3D) int {
	minDist := math.MaxInt

	distLeft := cube.X - min.X
	distRight := max.X - cube.X
	distUp := cube.Y - min.Y
	distDown := max.Y - cube.Y
	distBelow := cube.Z - min.Z
	distAbove := max.Z - cube.Z

	for _, val := range []int{distLeft, distRight, distUp, distDown, distBelow, distAbove} {
		if val < minDist {
			minDist = val
		}
	}

	return minDist
}

func isCubeInterior(cube Point3D, boundaryCubes s.Set[Point3D], min, max Point3D) (bool, s.Set[Point3D]) {

	queue := minpq.New[Point3D](0, 20*20*20)

	checked := s.NewSet(cube)

	priority := getPriority(cube, min, max)
	if priority < 0 {
		return false, checked
	}

	queue.AddAtPriority(cube, priority)

	// hacky: halt check if it takes too long
	for counter := 0; !queue.Empty() && counter <= 20; counter++ {
		toCheck := queue.PopMin()

		for _, neighbor := range toCheck.getAdjacent() {
			if !boundaryCubes.Contains(neighbor) && !checked.Contains(neighbor) {
				priority := getPriority(neighbor, min, max)

				if priority < 0 {
					checked.Add(neighbor)
					return false, checked
				} else {
					queue.AddAtPriority(neighbor, priority)
				}
			}
		}
	}

	return true, checked
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

func getMinMax(cubes []Point3D) (Point3D, Point3D) {
	min := Point3D{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Point3D{0, 0, 0}

	for _, cube := range cubes {
		if cube.X < min.X {
			min.X = cube.X
		}
		if cube.X > max.X {
			max.X = cube.X
		}
		if cube.Y < min.Y {
			min.Y = cube.Y
		}
		if cube.Y > max.Y {
			max.Y = cube.Y
		}
		if cube.Z < min.Z {
			min.Z = cube.Z
		}
		if cube.Z > max.Z {
			max.Z = cube.Z
		}
	}

	return min, max
}
