package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	point      geometry.Point
	spriteName string
	imageMap   map[string]*ebiten.Image
}

func NewUnit() Unit {
	return Unit{
		point:      geometry.NewPoint(0.0, 0.0),
		spriteName: "man",
		imageMap:   getUnitNameToImageMap(),
	}
}

func (u *Unit) DrawUnit(camera *render.Camera, screen *ebiten.Image) {
	imageToDraw := u.imageMap[u.spriteName]

	isoPoint := geometry.CartoToIso(u.point)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())

	camera.DrawImage(screen, imageToDraw, opts)
}

func getUnitNameToImageMap() map[string]*ebiten.Image {
	return map[string]*ebiten.Image{
		"man":   render.NewImageFromPath("./assets/units/Man.png"),
		"woman": render.NewImageFromPath("./assets/units/Woman.png"),
	}
}
