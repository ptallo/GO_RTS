package game

import (
	"go_rts/components"
	"go_rts/geometry"
	"go_rts/render"
)

// Unit is an object describing a game unit
type Unit struct {
	RenderComponent   components.IRenderComponent
	PositionComponent components.IPositionComponent
}

// NewUnit is a shorcut for creating a NewUnit
func NewUnit(ssl render.ISpriteSheetLibrary, camera render.ICamera) *Unit {
	u := Unit{
		RenderComponent:   components.NewRenderComponent(ssl, camera, "man"),
		PositionComponent: components.NewPositionComponent(geometry.NewPoint(0.0, 0.0), 5.0),
	}
	return &u
}
