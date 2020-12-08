package game

import (
	"go_rts/game/components"
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// Unit is an object describing a game unit
type Unit struct {
	spriteSheetLibrary render.ISpriteSheetLibrary
	camera             render.ICamera
	name               string
	PositionComponent  components.IPositionComponent
}

// NewUnit is a shorcut for creating a NewUnit
func NewUnit(ssl render.ISpriteSheetLibrary, camera render.ICamera) *Unit {
	u := Unit{
		spriteSheetLibrary: ssl,
		camera:             camera,
		name:               "man",
		PositionComponent:  components.NewPositionComponent(geometry.NewPoint(0.0, 0.0), 5.0),
	}
	return &u
}

// Draw is responsible for drawing a game unit on a screen
func (u *Unit) Draw(screen *ebiten.Image) {
	isoPoint := geometry.CartoToIso(*u.PositionComponent.GetPosition())
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(isoPoint.X, isoPoint.Y)
	ss := u.spriteSheetLibrary.GetSpriteSheet(u.name)
	ss.Draw(screen, u.camera, opts)
}

// GetDrawRectangle returns the rectangle of a unit in isometric space
func (u *Unit) GetDrawRectangle() geometry.Rectangle {
	isoPoint := geometry.CartoToIso(*u.PositionComponent.GetPosition())
	ss := u.spriteSheetLibrary.GetSpriteSheet(u.name)
	return geometry.NewRectangle(
		float64(ss.Definition.FrameWidth),
		float64(ss.Definition.FrameHeight),
		isoPoint.X,
		isoPoint.Y,
	)
}
