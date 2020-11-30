package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	spriteSheetLibrary *render.SpriteSheetLibrary
	camera             *render.Camera
	point              geometry.Point
	name               string
}

func NewUnit(ssl *render.SpriteSheetLibrary, camera *render.Camera) *Unit {
	u := Unit{
		spriteSheetLibrary: ssl,
		camera:             camera,
		point:              geometry.NewPoint(0.0, 0.0),
		name:               "man",
	}
	return &u
}

func (u *Unit) Draw(screen *ebiten.Image) {
	isoPoint := geometry.CartoToIso(u.point)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(isoPoint.X, isoPoint.Y)
	ss := u.spriteSheetLibrary.GetSpriteSheet(u.name)
	ss.Draw(screen, u.camera, opts)
}

func (u *Unit) GetDrawRectangle() geometry.Rectangle {
	isoPoint := geometry.CartoToIso(u.point)
	ss := u.spriteSheetLibrary.GetSpriteSheet(u.name)
	return geometry.NewRectangle(
		float64(ss.Definition.FrameWidth),
		float64(ss.Definition.FrameHeight),
		isoPoint.X,
		isoPoint.Y,
	)
}
