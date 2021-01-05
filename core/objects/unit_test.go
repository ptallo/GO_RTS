package objects_test

import (
	"go_rts/core/geometry"
	"go_rts/core/objects"
	"go_rts/core/render"
	"testing"
)

func Test_GivenIdenticalUnits_WhenCheckingEquality_ThenReturnsTrue(t *testing.T) {
	u1 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, 20.0, 20.0), 5.0),
	}

	u2 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, 20.0, 20.0), 5.0),
	}

	if !u1.Equals(u2) {
		t.Error("identical components when checking equality then should return true")
	}
}

func Test_GivenNonIdenticalUnits_WhenCheckingEquality_ThenReturnsFalse(t *testing.T) {
	u1 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, 20.0, 20.0), 5.0),
	}

	u2 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("woman"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, 20.0, 20.0), 5.0),
	}

	u3 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(22.0, 20.0, 20.0, 20.0), 5.0),
	}

	u4 := objects.Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, 20.0, 20.0), 6.0),
	}

	if u1.Equals(u2) || u1.Equals(u3) || u1.Equals(u4) {
		t.Error("non-identical components when checking equality then should return false")
	}
}
