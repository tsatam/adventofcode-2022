package main

import (
	c "common/cartesian"
	s "common/set"
)

type Rock struct {
	points s.Set[c.Point]
}

var Rocks = [5]Rock{
	{s.NewSet(
		/* ..####. */ c.Point{X: 2, Y: 0}, c.Point{X: 3, Y: 0}, c.Point{X: 4, Y: 0}, c.Point{X: 5, Y: 0},
	)},
	{s.NewSet(
		/* ...#... */ c.Point{X: 3, Y: -2},
		/* ..###.. */ c.Point{X: 2, Y: -1}, c.Point{X: 3, Y: -1}, c.Point{X: 4, Y: -1},
		/* ...#... */ c.Point{X: 3, Y: 0},
	)},
	{s.NewSet(
		/* ....#.. */ c.Point{X: 4, Y: -2},
		/* ....#.. */ c.Point{X: 4, Y: -1},
		/* ..###.. */ c.Point{X: 2, Y: 0}, c.Point{X: 3, Y: 0}, c.Point{X: 4, Y: 0},
	)},
	{s.NewSet(
		/* ..#.... */ c.Point{X: 2, Y: -3},
		/* ..#.... */ c.Point{X: 2, Y: -2},
		/* ..#.... */ c.Point{X: 2, Y: -1},
		/* ..#.... */ c.Point{X: 2, Y: -0},
	)},
	{s.NewSet(
		/* ..##... */ c.Point{X: 2, Y: -1}, c.Point{X: 3, Y: -1},
		/* ..##... */ c.Point{X: 2, Y: 0}, c.Point{X: 3, Y: 0},
	)},
}

// offsets are negative to make bigger numbers higher
func (r *Rock) AtOffset(yoffset int) Rock {
	newRock := Rock{s.NewSet[c.Point]()}

	for _, point := range r.points.Slice() {
		newRock.points.Add(c.Point{X: point.X, Y: point.Y - yoffset})
	}

	return newRock
}
