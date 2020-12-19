package geometry_test

import (
	"go_rts/geometry"
	"math"
	"testing"
)

func getMapPositionComponents() []geometry.IPositionComponent {
	tileWidth := 64.0
	tileHeight := 64.0
	pcs := make([]geometry.IPositionComponent, 0)
	for i := 0.0; i < 10.0; i++ {
		for j := 0.0; j < 10.0; j++ {
			rect := geometry.NewRectangle(tileWidth, tileHeight, i*tileWidth, j*tileHeight)
			pcs = append(pcs, geometry.NewPositionComponent(rect, 0.0))
		}
	}
	return pcs
}

func Test_WhenMovingTowardsDestination_ThenDistanceIsEffectedBySpeed(t *testing.T) {
	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 2.0
	p := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(dest, mapRect, []geometry.IPositionComponent{})
	p.MoveTowardsDestination([]geometry.IPositionComponent{})

	// Assert
	end := p.GetPosition()
	endDistance := math.Floor(end.DistanceFrom(dest)*1000) / 1000 // Rounded to not be too sensitive to floating point errors
	startDistance := math.Floor(start.DistanceFrom(dest)*1000) / 1000
	if endDistance != startDistance-speed {
		t.Errorf("end distance (%v) should equal start distance (%v) minus speed (%v)", endDistance, startDistance, speed)
	}
}

func Test_WhenMovingTowardsDesination_ThenWillNotOverStep(t *testing.T) {
	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 1000000.0
	p := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(dest, mapRect, []geometry.IPositionComponent{})
	p.MoveTowardsDestination([]geometry.IPositionComponent{})

	// Assert
	end := p.GetPosition()
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

func tryToMoveOutsideMap(t *testing.T, goalDestination geometry.Point, expectedDestination geometry.Point) {
	// Arrange
	p := geometry.NewPositionComponent(geometry.NewRectangle(5.0, 5.0, 200.0, 200.0), 3.0)
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	p.SetDestination(goalDestination, mapRect, []geometry.IPositionComponent{})

	// Assert
	actualDestination := *p.Destination
	if !expectedDestination.Equals(actualDestination) {
		t.Errorf("actual destination %v should equal expected destinaton %v", actualDestination, expectedDestination)
	}
}

func Test_GivenUnpathableComponent_WhenMoving_ThenCannotMoveThrough(t *testing.T) {
	// Arrange
	p1 := geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 0.0, 0.0), 5.0)
	pcs := []geometry.IPositionComponent{geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 0.0), 1000.0)}
	mapRect := geometry.NewRectangle(1000.0, 1000.0, 0.0, 0.0)

	// Act
	goalDestination := geometry.NewPoint(30.0, 0.0)
	p1.SetDestination(goalDestination, mapRect, pcs)

	for i := 0; i < 1000; i++ {
		p1.MoveTowardsDestination(pcs)
	}

	// Assert
	if p1.GetPosition().Equals(goalDestination) {
		t.Errorf("shouldn't be able to move through a position component between you and your desitination")
	}
}
