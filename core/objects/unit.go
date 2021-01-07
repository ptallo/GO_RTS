package objects

import (
	"go_rts/core/geometry"
	"go_rts/core/render"
)

// Unit is an object describing a game unit
type Unit struct {
	RenderComponent   *render.RenderComponent
	PositionComponent *geometry.PositionComponent
}

// NewUnit is a shorcut for creating a NewUnit
func NewUnit(startPosition geometry.Point) *Unit {
	u := Unit{
		RenderComponent:   render.NewRenderComponent("man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, startPosition.X, startPosition.Y), 0.5),
	}
	return &u
}

// Equals returns true if the two units are identical
func (u Unit) Equals(u2 Unit) bool {
	return u.RenderComponent.Equals(*u2.RenderComponent) && u.PositionComponent.Equals(*u2.PositionComponent)
}
