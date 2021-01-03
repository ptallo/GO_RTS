package render

import (
	"go_rts/types/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

// ICamera is the interface which is implemented to provide camera usage
type ICamera interface {
	DrawImage(*ebiten.Image, *ebiten.Image, *ebiten.DrawImageOptions)
	UpdateCameraPosition(float64, float64, geometry.Rectangle)
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

// UpdateCameraPosition will update the camera position according to GetCameraMovements if the screen will still overlap the map
func (c *Camera) UpdateCameraPosition(screenWidth, screenHeight float64, mapRect geometry.Rectangle) {
	moves := c.getCameraMovements()
	for _, move := range moves {
		c.translation.Translate(move)
	}

	screenRect := geometry.NewRectangle(
		screenWidth,
		screenHeight,
		c.translation.X,
		c.translation.Y,
	)

	if !screenRect.Intersects(mapRect) {
		for _, move := range moves {
			c.translation.Translate(move)
		}
	}
}

func (c *Camera) getCameraMovements() []geometry.Point {
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
