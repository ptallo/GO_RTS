package geometry_test

import (
	"go_rts/core/geometry"
	"math"
	"testing"
)

func Test_GivenIdenticalPositionComponents_WhenCheckingEquality_ThenReturnsTrue(t *testing.T) {
	p1 := geometry.NewPositionComponent(geometry.NewRectangle(11.1, 22.2, 33.3, 44.4), 55.5)
	p2 := *geometry.NewPositionComponent(geometry.NewRectangle(11.1, 22.2, 33.3, 44.4), 55.5)

	if !p1.Equals(p2) {
		t.Error("identical components when checking equality should return true")
	}
}

func Test_GivenNonIdenticalPositionComponents_WhenCheckingEquality_ThenReturnsFalse(t *testing.T) {
	p1 := geometry.NewPositionComponent(geometry.NewRectangle(11.1, 22.2, 33.3, 44.4), 55.5)
	p2 := *geometry.NewPositionComponent(geometry.NewRectangle(11.1, 22.2, 33.3, 66.6), 55.5)
	p3 := *geometry.NewPositionComponent(geometry.NewRectangle(11.1, 22.2, 33.3, 44.4), 66.6)

	if p1.Equals(p2) || p1.Equals(p3) {
		t.Error("non-identical components should return false when checking equality")
	}
}

func Test_WhenMovingTowardsDestination_ThenDistanceIsEffectedBySpeed(t *testing.T) {
	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 2.0
	p := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(dest, mapRect, []geometry.PositionComponent{})
	p.MoveTowardsDestination([]geometry.PositionComponent{})

	// Assert
	end := p.Rectangle.Point
	endDistance := math.Floor(end.DistanceFrom(dest)*1000) / 1000 // Rounded to not be too sensitive to floating point errors
	startDistance := math.Floor(start.DistanceFrom(dest)*1000) / 1000
	if endDistance != startDistance-speed {
		t.Errorf("end distance (%v) should equal start distance (%v) minus speed (%v)", endDistance, startDistance, speed)
	}
}

func Test_WhenMovingTowardsDestination_ThenWillNotOverStep(t *testing.T) {
	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 1000000.0
	p := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(dest, mapRect, []geometry.PositionComponent{})
	p.MoveTowardsDestination([]geometry.PositionComponent{})

	// Assert
	end := p.Rectangle.Point
	if end.DistanceFrom(dest) != 0.0 {
		t.Errorf("end should be ontop of the destination")
	}
}

func Test_GivenDestinationNotInMap_WhenSettingDestination_ThenStopsAtEdge(t *testing.T) {
	tryToMoveOutsideMap(t, geometry.NewPoint(-20.0, 200.0), geometry.NewPoint(0.0, 200.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(200.0, -20.0), geometry.NewPoint(200.0, 0.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(1200.0, 200.0), geometry.NewPoint(1000.0, 200.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(200.0, 1200.0), geometry.NewPoint(200.0, 1000.0))
}

func tryToMoveOutsideMap(t *testing.T, goalDestination, expectedDestination geometry.Point) {
	// Arrange
	p := geometry.NewPositionComponent(geometry.NewRectangle(5.0, 5.0, 200.0, 200.0), 3.0)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(goalDestination, mapRect, []geometry.PositionComponent{})

	// Assert
	actualDestination := p.GoalDestination
	if !expectedDestination.Equals(actualDestination) {
		t.Errorf("actual destination %v should equal expected destinaton %v", actualDestination, expectedDestination)
	}
}

func Test_GivenUnpathableComponent_WhenMoving_ThenCannotMoveThrough(t *testing.T) {
	// Arrange
	p1 := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 0.0, 0.0), 5.0)
	pc := *geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 0.0), 1000.0)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, -500.0, -500.0)

	// Act
	goalDestination := geometry.NewPoint(15.0, 0.0)
	p1.SetDestination(goalDestination, mapRect, []geometry.PositionComponent{pc})

	for i := 0; i < 1000; i++ {
		p1.MoveTowardsDestination([]geometry.PositionComponent{pc})
	}

	// Assert
	if p1.Rectangle.Point.Equals(goalDestination) || p1.Rectangle.Intersects(pc.Rectangle) {
		t.Error("shouldn't be able to move into un-pathable component")
	}
}

func Test_GivenUnpathableComponent_WhenMoving_ThenPathsAround(t *testing.T) {
	p1 := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 0.0, 0.0), 5.0)
	pcs := []geometry.PositionComponent{*geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 20.0, 0.0), 1000.0)}
	mapRect := geometry.NewRectangle(1000.0, 1000.0, -500.0, -500.0)

	// Act
	goalDestination := geometry.NewPoint(40.0, 0.0)
	p1.SetDestination(goalDestination, mapRect, pcs)

	for i := 0; i < 1000; i++ {
		p1.MoveTowardsDestination([]geometry.PositionComponent{})
	}

	// Assert
	if !p1.Rectangle.Point.Equals(goalDestination) {
		t.Error("Should path around un-pathable components")
	}
}
