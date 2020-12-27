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

func Test_GivenAdjacentRectangles_ThenIsAdjacentToReturnsTrue(t *testing.T) {
	// RightOf
	shouldBeAdjacent(
		geometry.NewRectangle(64.0, 64.0, 0.0, 0.0),
		geometry.NewRectangle(128.0, 64.0, 64.0, 0.0),
		t,
	)

	// LeftOf
	shouldBeAdjacent(
		geometry.NewRectangle(64.0, 64.0, 0.0, 0.0),
		geometry.NewRectangle(128.0, 64.0, -128.0, 0.0),
		t,
	)

	// Above
	shouldBeAdjacent(
		geometry.NewRectangle(64.0, 64.0, 0.0, 0.0),
		geometry.NewRectangle(64.0, 128.0, 0.0, -128.0),
		t,
	)

	// Below
	shouldBeAdjacent(
		geometry.NewRectangle(64.0, 64.0, 0.0, 0.0),
		geometry.NewRectangle(64.0, 128.0, 0.0, 64.0),
		t,
	)
}

func shouldBeAdjacent(r1, r2 geometry.Rectangle, t *testing.T) {
	if !r1.IsAdjacentTo(r2) {
		t.Errorf("rects %v and %v should be adjacent", r1.ToString(), r2.ToString())
	}
}

func Test_GivenNonAdjacentRectangles_ThenIsAdjacentToReturnsFalse(t *testing.T) {
	shouldNotBeAdjacent(
		geometry.NewRectangle(10.0, 10.0, 0.0, 0.0),
		geometry.NewRectangle(10.0, 10.0, 10.0, 10.0),
		t,
	)

	shouldNotBeAdjacent(
		geometry.NewRectangle(10.0, 10.0, 0.0, 0.0),
		geometry.NewRectangle(10.0, 10.0, -10.0, -10.0),
		t,
	)

	shouldNotBeAdjacent(
		geometry.NewRectangle(10.0, 10.0, 0.0, 0.0),
		geometry.NewRectangle(10.0, 10.0, 10.0, -10.0),
		t,
	)

	shouldNotBeAdjacent(
		geometry.NewRectangle(10.0, 10.0, 0.0, 0.0),
		geometry.NewRectangle(10.0, 10.0, -10.0, 10.0),
		t,
	)
}

func shouldNotBeAdjacent(r1, r2 geometry.Rectangle, t *testing.T) {
	if r1.IsAdjacentTo(r2) {
		t.Errorf("rects %v and %v should not be adjacent", r1.ToString(), r2.ToString())
	}
}
