package game

import (
	"go_rts/game/components"
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
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

// Draw is responsible for drawing a game unit on a screen
func (u *Unit) Draw(screen *ebiten.Image) {
	isoPoint := geometry.CartoToIso(*u.PositionComponent.GetPosition())
	u.RenderComponent.Draw(screen, isoPoint)
}

// GetDrawRectangle returns the rectangle of a unit in isometric space
func (u *Unit) GetDrawRectangle() geometry.Rectangle {
	isoPoint := geometry.CartoToIso(*u.PositionComponent.GetPosition())
	return u.RenderComponent.GetDrawRectangle(isoPoint)
}
