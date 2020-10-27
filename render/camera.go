package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	translation Point
	speed       float64
}

func NewCamera() Camera {
	return Camera{
		translation: NewPoint(0, 0),
		speed:       5.0,
	}
}

func (c *Camera) DrawImage(screen, img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(c.translation.x, c.translation.y)
	screen.DrawImage(img, opts)
}

func (c *Camera) MoveCamera(p Point) {
	c.translation.Move(p)
}

func (c *Camera) UpdateCameraPosition() {
	width, height := ebiten.WindowSize()
	cursorX, cursorY := ebiten.CursorPosition()
	p := NewPoint(float64(cursorX), float64(cursorY))
	if p.X() < float64(width)*0.1 {
		c.MoveCamera(NewPoint(c.speed, 0))
	}
	if p.X() > float64(width)*0.9 {
		c.MoveCamera(NewPoint(-c.speed, 0))
	}
	if p.Y() < float64(height)*0.1 {
		c.MoveCamera(NewPoint(0, c.speed))
	}
	if p.Y() > float64(height)*0.9 {
		c.MoveCamera(NewPoint(0, -c.speed))
	}
}
