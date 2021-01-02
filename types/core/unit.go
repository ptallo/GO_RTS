package core

import (
	"go_rts/client/render"
	"go_rts/types/geometry"
)

// Unit is an object describing a game unit
type Unit struct {
	RenderComponent   render.IRenderComponent
	PositionComponent geometry.IPositionComponent
}

// NewUnit is a shorcut for creating a NewUnit
func NewUnit(ssl render.ISpriteSheetLibrary, camera render.ICamera, startPosition geometry.Point) *Unit {
	u := Unit{
		RenderComponent:   render.NewRenderComponent(ssl, camera, "man"),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(20.0, 20.0, startPosition.X, startPosition.Y), 5.0),
	}
	return &u
}
