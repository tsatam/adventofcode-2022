package cartesian

type Point struct{ X, Y int }

func (p Point) Move(direction Direction) Point {
	switch direction {
	case Up:
		return Point{p.X, p.Y - 1}
	case Down:
		return Point{p.X, p.Y + 1}
	case Left:
		return Point{p.X - 1, p.Y}
	case Right:
		return Point{p.X + 1, p.Y}
	default:
		return p
	}
}
