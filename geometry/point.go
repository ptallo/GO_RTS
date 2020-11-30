package geometry

type Point struct {
	X float64
	Y float64
}

func NewPoint(newX, newY float64) Point {
	return Point{
		X: newX,
		Y: newY,
	}
}

func (p *Point) Translate(point Point) {
	p.X += point.X
	p.Y += point.Y
}

func (p Point) Inverse() Point {
	return NewPoint(-p.X, -p.Y)
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func CartoToIso(p Point) Point {
	return Point{
		X: p.X - p.Y,
		Y: (p.X + p.Y) / 2,
	}
}

func IsoToCarto(p Point) Point {
	return Point{
		X: (2*p.Y + p.X) / 2,
		Y: (2*p.Y - p.X) / 2,
	}
}
