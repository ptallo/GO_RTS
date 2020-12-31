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
	return r.IsTopAdjacent(r2) ||
		r.IsBottomAdjacent(r2) ||
		r.IsLeftAdjacent(r2) ||
		r.IsRightAdjacent(r2)
}

// IsTopAdjacent returns true if r is above r2 else false
func (r Rectangle) IsTopAdjacent(r2 Rectangle) bool {
	return r.Point.Y+r.Height == r2.Point.Y && r.Point.X == r2.Point.X && r.Width == r2.Width
}

// IsBottomAdjacent returns true if r is below r2 else false
func (r Rectangle) IsBottomAdjacent(r2 Rectangle) bool {
	return r.Point.Y == r2.Point.Y+r2.Height && r.Point.X == r2.Point.X && r.Width == r2.Width
}

// IsLeftAdjacent returns true if r is left of r2 else false
func (r Rectangle) IsLeftAdjacent(r2 Rectangle) bool {
	return r.Point.X+r.Width == r2.Point.X && r.Point.Y == r2.Point.Y && r.Height == r2.Height
}

// IsRightAdjacent returns true if r is right of r2 else false
func (r Rectangle) IsRightAdjacent(r2 Rectangle) bool {
	return r.Point.X == r2.Point.X+r2.Width && r.Point.Y == r2.Point.Y && r.Height == r2.Height
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

// Center returns the point describing the center of the rectangle
func (r Rectangle) Center() Point {
	return NewPoint(
		r.Point.X+r.Width/2,
		r.Point.Y+r.Height/2,
	)
}

// ToString prints a rectangle to a string in a nice way
func (r Rectangle) ToString() string {
	return fmt.Sprintf("%g %g %v", r.Width, r.Height, r.Point)
}
