package geometry

// Point defines a single point in a 2d plane
type Point struct {
	X float64
	Y float64
}

// NewPoint is a shortcut to create a new point given a x and y value
func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Translate moves a point given another point
func (p *Point) Translate(point Point) {
	p.X += point.X
	p.Y += point.Y
}

// Inverse gives the inverse of a given point
func (p Point) Inverse() Point {
	return NewPoint(-p.X, -p.Y)
}

// Equals checks if two points are equal
func (p Point) Equals(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

// CartoToIso converts a point from a Cartesian point to a Isometric point
func CartoToIso(p Point) Point {
	return Point{
		X: p.X - p.Y,
		Y: (p.X + p.Y) / 2,
	}
}

// IsoToCarto converts a point from a Isometric point to a Isometric point
func IsoToCarto(p Point) Point {
	return Point{
		X: (2*p.Y + p.X) / 2,
		Y: (2*p.Y - p.X) / 2,
	}
}
