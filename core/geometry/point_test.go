package geometry_test

import (
	"go_rts/core/geometry"
	"testing"
)

func Test_GivenPoint_AndDestination_ThenCanGetPathToDestination(t *testing.T) {
	p1 := geometry.NewPoint(5.0, 0.0)
	p2 := geometry.NewPoint(5.0, 5.0)

	if !p1.To(p2).Equals(geometry.NewPoint(0.0, 5.0)) {
		t.Errorf("expected to recieve point to p2")
	}
}

func Test_GivenPoint_WhenCallingUnit_ReturnsScaledPointOfLengthOne(t *testing.T) {
	p1 := geometry.NewPoint(10.0, 10.0)

	unit := p1.Unit()

	expected := geometry.NewPoint(0.7071067811865475, 0.7071067811865475)
	if !unit.Equals(expected) {
		t.Errorf("expected normalized vector to equal %v but was %v", expected, unit)
	}
}
