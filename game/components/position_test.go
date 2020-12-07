package components_test

import (
	"go_rts/game/components"
	"go_rts/geometry"
	"math"
	"testing"
)

func Test_WhenMovingTowardsDestination_ThenDistanceIsEffectedBySpeed(t *testing.T) {
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 2.0
	p := components.NewPositionComponent(start, speed)
	p.SetDestination(dest)

	p.MoveTowardsDestination()
	end := p.GetPosition()

	endDistance := math.Floor(end.DistanceFrom(dest)*1000) / 1000 // Rounded to not be too sensitive to floating point errors
	startDistance := math.Floor(start.DistanceFrom(dest)*1000) / 1000
	if endDistance != startDistance-speed {
		t.Errorf("end distance (%v) should equal start distance (%v) minus speed (%v)", endDistance, startDistance, speed)
	}
}

func Test_WhenMovingTowardsDesination_ThenWillNotOverStep(t *testing.T) {
	start := geometry.NewPoint(0.0, 0.0)
	dest := geometry.NewPoint(10.0, 10.0)
	speed := 1000000.0
	p := components.NewPositionComponent(start, speed)
	p.SetDestination(dest)

	p.MoveTowardsDestination()
	end := p.GetPosition()

	if end.DistanceFrom(dest) != 0.0 {
		t.Errorf("end should be ontop of the destination")
	}
}
