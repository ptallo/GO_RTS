package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	point geometry.Point
	name  string
}

func NewUnit() *Unit {
	u := Unit{
		point: geometry.NewPoint(0.0, 0.0),
		name:  "man",
	}
	return &u
}

func (u *Unit) Draw(camera *render.Camera, screen *ebiten.Image, lib map[string]*render.SpriteSheet) {
	isoPoint := geometry.CartoToIso(u.point)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())
	spriteSheet := lib[u.name]
	spriteSheet.Draw(screen, camera, opts)
}
