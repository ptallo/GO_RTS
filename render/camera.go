package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	translation Point
}

func NewCamera() Camera {
	return Camera{
		translation: NewPoint(0, 0),
	}
}

func (c *Camera) DrawImage(screen, img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(c.translation.X, c.translation.Y)
	screen.DrawImage(img, opts)
}

func (c *Camera) MoveCamera(p Point) {
	c.translation.Move(p)
}
