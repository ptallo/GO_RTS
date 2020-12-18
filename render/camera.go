package render

import (
	"go_rts/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)


// ICamera is the interface which is implemented to provide camera usage
type ICamera interface {
	DrawImage(*ebiten.Image, *ebiten.Image, *ebiten.DrawImageOptions)
	GetCameraMovements() []geometry.Point
	Translation() *geometry.Point
}

// Camera is an object used to track how the screen has moved in a game
type Camera struct {
	translation *geometry.Point
	Speed       float64
}

// NewCamera is a shortcut to creating a Camera object
func NewCamera(t *geometry.Point, s float64) ICamera {
	return &Camera{
		translation: t,
		Speed:       s,
	}
}

// DrawImage draws an image on a screen adjusted for the cameras translation
func (c *Camera) DrawImage(screen, img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(-c.translation.X, -c.translation.Y)
	screen.DrawImage(img, opts)
}

// GetCameraMovements will return how the camera should move given the current position of a mouse in relation to the window
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

// Translation gives the cameras current position
func (c *Camera) Translation() *geometry.Point {
	return c.translation
}
