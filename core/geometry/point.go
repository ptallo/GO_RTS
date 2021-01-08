package geometry

import (
	"math"
)

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

// Move returns a new point moved by the value p
func (p Point) Move(p1 Point) Point {
	return NewPoint(p.X+p1.X, p.Y+p1.Y)
}

// Inverse gives the inverse of a given point
func (p Point) Inverse() Point {
	return NewPoint(-p.X, -p.Y)
}

// Equals checks if two points are equal
func (p Point) Equals(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

// DistanceFrom returns the distance from this point to another
func (p Point) DistanceFrom(point Point) float64 {
	return math.Sqrt(math.Pow(p.X-point.X, 2) + math.Pow(p.Y-point.Y, 2))
}

// To return a vector describing the point from p to the argument
func (p Point) To(point Point) Point {
	return Point{
		X: point.X - p.X,
		Y: point.Y - p.Y,
	}
}

// Unit returns the current point scaled to unit length
func (p Point) Unit() Point {
	length := p.DistanceFrom(NewPoint(0.0, 0.0))
	return Point{
		X: p.X / length,
		Y: p.Y / length,
	}
}

// Scale scales the current vector by a certain scale
func (p Point) Scale(scale float64) Point {
	return Point{
		X: p.X * scale,
		Y: p.Y * scale,
	}
}
