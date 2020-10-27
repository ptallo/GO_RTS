package render

type Point struct {
	x float64
	y float64
}

func NewPoint(newX, newY float64) Point {
	return Point{
		x: newX,
		y: newY,
	}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p *Point) Move(point Point) {
	p.x += point.x
	p.y += point.y
}

func CartoToIso(p Point) Point {
	return Point{
		x: p.x - p.y,
		y: (p.x + p.y) / 2,
	}
}

func IsoToCarto(p Point) Point {
	return Point{
		x: (2*p.x + p.y) / 2,
		y: (2 + p.x - p.y) / 2,
	}
}
