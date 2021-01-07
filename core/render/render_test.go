package render_test

import (
	"go_rts/core/render"
	"testing"
)

func Test_GivenIdenticalComponent_WhenCheckingEquality_ThenReturnsTrue(t *testing.T) {
	r1 := render.NewRenderComponent("test")
	r2 := render.NewRenderComponent("test")

	if !r1.Equals(r2) {
		t.Error("identical render components should return true when checking equality")
	}
}

func Test_GivenNonIdenticalComponents_WhenCheckingEquality_ThenReturnsFalse(t *testing.T) {
	r1 := render.NewRenderComponent("test")
	r2 := render.NewRenderComponent("test2")

	if r1.Equals(r2) {
		t.Error("identical render components should return true when checking equality")
	}
}
