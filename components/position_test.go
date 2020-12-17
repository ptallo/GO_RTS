package components_test

import (
	"go_rts/components"
	"go_rts/game"
	"go_rts/geometry"
	"go_rts/mocks"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
)

func getMapPositionComponents(ctrl *gomock.Controller) []components.IPositionComponent {
	mockSSL := mocks.NewMockISpriteSheetLibrary(ctrl)
	mockCamera := mocks.NewMockICamera(ctrl)

	tiles := game.NewMap(mockSSL, mockCamera)
	pcs := make([]components.IPositionComponent, 0)
	for _, tile := range tiles {
		pcs = append(pcs, tile.PositionComponent)
	}
	return pcs
}

func Test_WhenMovingTowardsDestination_ThenDistanceIsEffectedBySpeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 2.0
	p := components.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	pcs := getMapPositionComponents(ctrl)

	// Act
	p.SetDestination(dest, pcs)
	p.MoveTowardsDestination([]components.IPositionComponent{})

	// Assert
	end := p.GetPosition()
	endDistance := math.Floor(end.DistanceFrom(dest)*1000) / 1000 // Rounded to not be too sensitive to floating point errors
	startDistance := math.Floor(start.DistanceFrom(dest)*1000) / 1000
	if endDistance != startDistance-speed {
		t.Errorf("end distance (%v) should equal start distance (%v) minus speed (%v)", endDistance, startDistance, speed)
	}
}

func Test_WhenMovingTowardsDesination_ThenWillNotOverStep(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 1000000.0
	p := components.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, start.X, start.Y), speed)
	pcs := getMapPositionComponents(ctrl)

	// Act
	p.SetDestination(dest, pcs)
	p.MoveTowardsDestination([]components.IPositionComponent{})

	// Assert
	end := p.GetPosition()
	if end.DistanceFrom(dest) != 0.0 {
		t.Errorf("end should be ontop of the destination")
	}
}

func Test_GivenDestinationNotInMap_WhenSettingDestination_ThenStopsAtEdge(t *testing.T) {
	tryToMoveOutsideMap(t, geometry.NewPoint(-20.0, 200.0), geometry.NewPoint(0.0, 192.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(200.0, -20.0), geometry.NewPoint(192.0, 0.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(1000.0, 200.0), geometry.NewPoint(640.0, 192.0))
	tryToMoveOutsideMap(t, geometry.NewPoint(200.0, 1000.0), geometry.NewPoint(192.0, 640.0))
}

func tryToMoveOutsideMap(t *testing.T, goalDestination geometry.Point, expectedDestination geometry.Point) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	pcs := getMapPositionComponents(ctrl)
	p := components.NewPositionComponent(geometry.NewRectangle(5.0, 5.0, 200.0, 200.0), 3.0)

	// Act
	p.SetDestination(goalDestination, pcs)

	// Assert
	actualDestination := *p.Destination
	if !expectedDestination.Equals(actualDestination) {
		t.Errorf("actual destination %v should equal expected destinaton %v", actualDestination, expectedDestination)
	}
}

func Test_GivenUnpathableComponent_WhenMoving_ThenCannotMoveThrough(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	tiles := getMapPositionComponents(ctrl)
	p1 := components.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 0.0, 0.0), 5.0)
	pcs := []components.IPositionComponent{components.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 0.0), 1000.0)}

	// Act
	goalDestination := geometry.NewPoint(30.0, 0.0)
	p1.SetDestination(goalDestination, tiles)
	for i := 0; i < 1000; i++ {
		p1.MoveTowardsDestination(pcs)
	}

	// Assert
	if p1.GetPosition().Equals(goalDestination) {
		t.Errorf("shouldn't be able to move through a position component between you and your desitination")
	}
}
