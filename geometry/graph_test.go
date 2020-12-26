package geometry_test

import (
	"go_rts/geometry"
	"testing"
)

func Test_GivenPositionComponents_ThenGeneratesGraph(t *testing.T) {
	// Arrange
	pcs := []geometry.IPositionComponent{
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 0.0), 0.0),
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 64.0), 0.0),
		geometry.NewPositionComponent(geometry.NewRectangle(64.0, 64.0, 128.0, 128.0), 0.0),
	}
	mapRect := geometry.NewRectangle(256.0, 256.0, 0.0, 0.0)

	// Act
	graph := geometry.NewGraph(pcs, mapRect)

	// Assert
	if len(graph.AdjacencyList) != 9 {
		t.Errorf("should construct a graph with 9 nodes")
	}

	for k, v := range graph.AdjacencyList {
		checkRectInCollidableComponents(k, pcs, t)

		if len(v) != 1 && len(v) != 2 {
			t.Errorf("all elements should have 1 or 2 adjacent nodes")
		}
	}
}

func checkRectInCollidableComponents(rect geometry.Rectangle, pcs []geometry.IPositionComponent, t *testing.T) {
	for _, pc := range pcs {
		if rect.Equals(pc.GetRectangle()) {
			t.Errorf("Collidable components should not appear in map")
		}
	}
}
