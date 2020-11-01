package geometry_test

import (
	"go_rts/geometry"
	"testing"
)

func Test_GivenRectangle_WhenContainsPoint_ThenReturnsTrue(t *testing.T) {
	rect := geometry.NewRectangle(10.0, 10.0, 0.0, 0.0)
	p := geometry.NewPoint(5.0, 5.0)

	if !rect.Contains(p) {
		t.Errorf("rect: %v should contain point: %v but doesn't", rect, p)
	}
}

func Test_GivenRectangle_WhenPointOnEdge_ThenReturnsTrue(t *testing.T) {
	rect := geometry.NewRectangle(10.0, 10.0, 0.0, 0.0)
	points := []geometry.Point{
		geometry.NewPoint(0.0, 0.0),
		geometry.NewPoint(10.0, 0.0),
		geometry.NewPoint(0.0, 10.0),
		geometry.NewPoint(10.0, 10.0),
	}

	for _, p := range points {
		if !rect.Contains(p) {
			t.Errorf("rect: %v should contain point: %v but doesn't", rect, p)
		}
	}
}

func Test_GivenRectangle_WhenDoesntContainPoint_ThenReturnsFalse(t *testing.T) {
	rects := []geometry.Rectangle{
		geometry.NewRectangle(10.0, 5.0, 0.0, 0.0),
		geometry.NewRectangle(5.0, 10.0, 0.0, 0.0),
		geometry.NewRectangle(10.0, 5.0, 11.0, 11.0),
		geometry.NewRectangle(5.0, 10.0, 11.0, 11.0),
	}

	p := geometry.NewPoint(10.0, 10.0)

	for _, rect := range rects {
		if rect.Contains(p) {
			t.Errorf("rect: %v shouldn't contain point: %v but does", rect, p)
		}
	}
}
