package geometry

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
	Point  *Point
}

func NewRectangle(w, h, x, y float64) Rectangle {
	p := NewPoint(x, y)
	return Rectangle{
		Width:  w,
		Height: h,
		Point:  &p,
	}
}

func NewRectangleFromPoints(p1, p2 Point) Rectangle {
	x := math.Min(p1.X, p2.X)
	x2 := math.Max(p1.X, p2.X)
	y := math.Min(p1.Y, p2.Y)
	y2 := math.Max(p1.Y, p2.Y)

	return NewRectangle(x2-x, y2-y, x, y)
}

func (r Rectangle) Translate(p Point) {
	r.Point.Translate(p)
}

func (r1 Rectangle) Intersects(r2 Rectangle) bool {
	if r1.Point.X >= r2.Point.X+r2.Width || r2.Point.X >= r1.Point.X+r1.Width {
		return false
	}
	if r1.Point.Y >= r2.Point.Y+r2.Height || r2.Point.Y >= r1.Point.Y+r1.Height {
		return false
	}
	return true
}

func (r Rectangle) Contains(p Point) bool {
	return r.Point.X <= p.X && r.Point.X+r.Width >= p.X && r.Point.Y <= p.Y && r.Point.Y+r.Height >= p.Y
}

func (r Rectangle) ToString() string {
	return fmt.Sprintf("%g %g %v", r.Width, r.Height, r.Point)
}
