package game

type Point struct {
	x float64
	y float64
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
