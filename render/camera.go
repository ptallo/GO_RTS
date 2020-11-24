package render

import (
	"go_rts/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	translation geometry.Point
	speed       float64
}

func NewCamera() *Camera {
	c := Camera{
		translation: geometry.NewPoint(0, 0),
		speed:       5.0,
	}
	return &c
}

func (c *Camera) Translation() geometry.Point {
	return c.translation
}

func (c *Camera) Speed() float64 {
	return c.speed
}

func (c *Camera) DrawImage(screen, img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(-c.translation.X(), -c.translation.Y())
	screen.DrawImage(img, opts)
}

func (c *Camera) MoveCamera(p geometry.Point) {
	c.translation.Move(p)
}

func (c *Camera) GetCameraMovements() []geometry.Point {
	width, height := ebiten.WindowSize()

	cursorX, cursorY := ebiten.CursorPosition()
	p := geometry.NewPoint(float64(cursorX), float64(cursorY))

	moves := make([]geometry.Point, 0)
	if p.X() < float64(width)*0.1 {
		moves = append(moves, geometry.NewPoint(-c.Speed(), 0))
	}
	if p.X() > float64(width)*0.9 {
		moves = append(moves, geometry.NewPoint(c.Speed(), 0))
	}
	if p.Y() < float64(height)*0.1 {
		moves = append(moves, geometry.NewPoint(0, -c.Speed()))
	}
	if p.Y() > float64(height)*0.9 {
		moves = append(moves, geometry.NewPoint(0, c.Speed()))
	}

	return moves
}
