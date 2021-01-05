package objects_test

import (
	"go_rts/core/geometry"
	"go_rts/core/objects"
	"go_rts/core/render"
	"testing"
)

func Test_GivenIdenticalTiles_WhenCheckingEquality_ThenReturnsTrue(t *testing.T) {
	tile1 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	tile2 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	if !tile1.Equals(tile2) {
		t.Error("identical tiles should return true when checking for equality")
	}
}

func Test_GivenNonIdenticalTiles_WhenCheckingEquality_ThenReturnsFalse(t *testing.T) {
	tile1 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	tile2 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent1"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(10.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	tile3 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent1"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(100.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	tile4 := objects.Tile{
		RenderComponent:   render.NewRenderComponent("testcomponent1"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(100.0, 10.0, 10.0, 10.0), 0.0),
		IsPathable:        false,
	}

	if tile1.Equals(tile2) || tile1.Equals(tile3) || tile1.Equals(tile4) {
		t.Error("non-idnetical tiles should return false when checking for equality")
	}
}
