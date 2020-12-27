package geometry_test

import (
	"go_rts/geometry"
	"testing"
)

func Test_GivenNoCollidables_ThenReturnsDest(t *testing.T) {
	graph := geometry.NewGraph([]geometry.IPositionComponent{}, geometry.NewRectangle(256.0, 256.0, 0.0, 0.0))
	start := geometry.NewPoint(10.0, 10.0)
	destination := geometry.NewPoint(100.0, 100.0)

	path := graph.PathFrom(start, destination)

	if len(path) != 1 && !path[0].Contains(destination) {
		t.Errorf("path should be length 1 and have goal destination %v but instead has destination %v", destination, path[0])
	}
}

func Test_GivenGraph_WhenPathing_ThenGeneratesPoints(t *testing.T) {
	// Arrange
	pcs := []geometry.IPositionComponent{
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 0.0), 0.0),
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 64.0), 0.0),
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 128.0), 0.0),
	}
	mapRect := geometry.NewRectangle(256.0, 256.0, 0.0, 0.0)
	graph := geometry.NewGraph(pcs, mapRect)

	start := geometry.NewPoint(10.0, 10.0)
	dest := geometry.NewPoint(200.0, 10.0)

	// Act
	nodesToDestination := graph.PathFrom(start, dest)

	// Assert
	if len(nodesToDestination) != 9 {
		t.Errorf("number of nodes to destination should be 8 but is %v", len(nodesToDestination))
	}

	lastNode := nodesToDestination[len(nodesToDestination)-1]
	if !lastNode.Contains(dest) {
		t.Error("The last node should contain the destination")
	}

	startNode := nodesToDestination[0]
	if !startNode.Contains(start) {
		t.Error("The first node should contain the start")
	}
}
