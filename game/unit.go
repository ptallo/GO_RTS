package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	point      geometry.Point
	spriteName string
}

func NewUnit() Unit {
	return Unit{
		point:      geometry.NewPoint(0.0, 0.0),
		spriteName: "man",
	}
}

func (u *Unit) DrawUnit(camera *render.Camera, screen *ebiten.Image, imageLib *render.ImageLibrary) {
	imageToDraw, err := imageLib.GetImage(u.spriteName)

	if err != nil {
		panic(err)
	}

	isoPoint := geometry.CartoToIso(u.point)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())

	camera.DrawImage(screen, imageToDraw, opts)
}
