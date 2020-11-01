package geometry

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
	point  *Point
}

func NewRectangle(w, h, x, y float64) Rectangle {
	p := NewPoint(x, y)
	return Rectangle{
		width:  w,
		height: h,
		point:  &p,
	}
}

func (r Rectangle) Move(x, y float64) {
	r.point.Move(NewPoint(x, y))
}

func (r1 Rectangle) Intersects(r2 Rectangle) bool {
	if r1.point.x >= r2.point.x+r2.width || r2.point.x >= r1.point.x+r1.width {
		return false
	}
	if r1.point.y >= r2.point.y+r2.height || r2.point.y >= r1.point.y+r1.height {
		return false
	}
	return true
}

func (r Rectangle) Contains(p Point) bool {
	return r.point.x <= p.x && r.point.x+r.width >= p.x && r.point.y <= p.y && r.point.y+r.height >= p.y
}

func (r Rectangle) ToString() string {
	return fmt.Sprintf("%g %g %v", r.width, r.height, r.point)
}
