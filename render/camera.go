package render

import (
	"go_rts/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Translation geometry.Point
	Speed       float64
}

func NewCamera(t geometry.Point, s float64) *Camera {
	return &Camera{
		Translation: t,
		Speed:       s,
	}
}

func (c *Camera) DrawImage(screen, img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(-c.Translation.X, -c.Translation.Y)
	screen.DrawImage(img, opts)
}

func (c *Camera) MoveCamera(p geometry.Point) {
	c.Translation.Translate(p)
}

func (c *Camera) GetCameraMovements() []geometry.Point {
	width, height := ebiten.WindowSize()

	cursorX, cursorY := ebiten.CursorPosition()
	p := geometry.NewPoint(float64(cursorX), float64(cursorY))

	moves := make([]geometry.Point, 0)
	if p.X < float64(width)*0.1 {
		moves = append(moves, geometry.NewPoint(-c.Speed, 0))
	}
	if p.X > float64(width)*0.9 {
		moves = append(moves, geometry.NewPoint(c.Speed, 0))
	}
	if p.Y < float64(height)*0.1 {
		moves = append(moves, geometry.NewPoint(0, -c.Speed))
	}
	if p.Y > float64(height)*0.9 {
		moves = append(moves, geometry.NewPoint(0, c.Speed))
	}

	return moves
}
