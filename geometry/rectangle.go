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

func (r1 Rectangle) IsOverlapping(r2 Rectangle) bool {
	if r1.point.x >= r2.point.x+r2.width || r2.point.x >= r1.point.x+r1.width {
		return false
	}
	if r1.point.y >= r2.point.y+r2.height || r2.point.y >= r1.point.y+r1.height {
		return false
	}
	return true
}

func (r Rectangle) ToString() string {
	return fmt.Sprintf("%g %g %v", r.width, r.height, r.point)
}
