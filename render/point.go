package render

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

func (p *Point) Move(point Point) {
	p.X += point.X
	p.Y += point.Y
}

func CartoToIso(p Point) Point {
	return Point{
		X: p.X - p.Y,
		Y: (p.X + p.Y) / 2,
	}
}

func IsoToCarto(p Point) Point {
	return Point{
		X: (2*p.X + p.Y) / 2,
		Y: (2 + p.X - p.Y) / 2,
	}
}
