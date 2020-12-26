package geometry

import (
	"fmt"
	"math"
)

// Rectangle is defined given its top-left point and the width and height
type Rectangle struct {
	Width  float64
	Height float64
	Point  *Point
}

// NewRectangle is a shortcut to define a new rectangle
func NewRectangle(w, h, x, y float64) Rectangle {
	p := NewPoint(x, y)
	return Rectangle{
		Width:  w,
		Height: h,
		Point:  &p,
	}
}

// NewRectangleFromPoints is a shortcut to define a new Rectangle given two points
func NewRectangleFromPoints(p1, p2 Point) Rectangle {
	x := math.Min(p1.X, p2.X)
	x2 := math.Max(p1.X, p2.X)
	y := math.Min(p1.Y, p2.Y)
	y2 := math.Max(p1.Y, p2.Y)

	return NewRectangle(x2-x, y2-y, x, y)
}

// Equals checks to see if this rect is equal to another
func (r Rectangle) Equals(r2 Rectangle) bool {
	return r.Point.Equals(*r2.Point) && r.Width == r2.Width && r.Height == r.Height
}

// IsAdjacentTo checkts to see if this rect exists next to another in 2d space
func (r Rectangle) IsAdjacentTo(r2 Rectangle) bool {
	return r.isAbove(r2) ||
		r.isBelow(r2) ||
		r.isLeftOf(r2) ||
		r.isRightOf(r2)
}

func (r Rectangle) isAbove(r2 Rectangle) bool {
	return r.Point.Y+r.Height == r2.Point.Y && r.Point.X == r2.Point.X && r.Width == r2.Width
}

func (r Rectangle) isBelow(r2 Rectangle) bool {
	return r.Point.Y == r2.Point.Y+r2.Height && r.Point.X == r2.Point.X && r.Width == r2.Width
}

func (r Rectangle) isLeftOf(r2 Rectangle) bool {
	return r.Point.X+r.Width == r2.Point.X && r.Point.Y == r2.Point.Y && r.Width == r2.Width
}

func (r Rectangle) isRightOf(r2 Rectangle) bool {
	return r.Point.X == r2.Point.X+r2.Width && r.Point.Y == r2.Point.Y && r.Width == r2.Width
}

// Intersects checks if one rectangle intersects another rectangle
func (r Rectangle) Intersects(r2 Rectangle) bool {
	if r.Point.X >= r2.Point.X+r2.Width || r2.Point.X >= r.Point.X+r.Width {
		return false
	}
	if r.Point.Y >= r2.Point.Y+r2.Height || r2.Point.Y >= r.Point.Y+r.Height {
		return false
	}
	return true
}

// Contains checks to see if a rectangle contains a point
func (r Rectangle) Contains(p Point) bool {
	return r.Point.X <= p.X && r.Point.X+r.Width >= p.X && r.Point.Y <= p.Y && r.Point.Y+r.Height >= p.Y
}

// ToString prints a rectangle to a string in a nice way
func (r Rectangle) ToString() string {
	return fmt.Sprintf("%g %g %v", r.Width, r.Height, r.Point)
}

// GetCorners returns the four-corners of a rectangle
func (r Rectangle) GetCorners() []Point {
	x := r.Point.X
	y := r.Point.Y
	return []Point{
		NewPoint(x, y),
		NewPoint(x+r.Width, y),
		NewPoint(x, y+r.Height),
		NewPoint(x+r.Width, y+r.Height),
	}
}
