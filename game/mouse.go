package game

import (
	"go_rts/geometry"
	"go_rts/render"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mouse struct {
	leftButtonDownDuration int
	leftButtonDownPoint    geometry.Point
	spriteSheet            *render.SpriteSheet
}

func NewMouse() *Mouse {
	m := Mouse{
		leftButtonDownDuration: 0,
		leftButtonDownPoint:    geometry.NewPoint(0, 0),
	}
	return &m
}

func (m *Mouse) Update() {
	m.updateLeftButtonPressedDuration()

	if m.isLeftButtonJustPressed() {
		m.leftButtonDownPoint = m.position()
	}
}

func (m *Mouse) updateLeftButtonPressedDuration() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		m.leftButtonDownDuration++
	} else {
		m.leftButtonDownDuration = 0
	}
}

func (m Mouse) isLeftButtonJustPressed() bool {
	return m.leftButtonDownDuration == 1
}

func (m Mouse) isLeftButtonPressed() bool {
	return m.leftButtonDownDuration != 0
}

func (m *Mouse) Draw(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		rect := m.getMouseSelectionRect()
		opts := m.getMouseDrawOptions(rect)
		img := m.getMouseImage(int(rect.Width()), int(rect.Height()))
		screen.DrawImage(img, opts)
	} else {
	}
}

func (m *Mouse) getMouseSelectionRect() geometry.Rectangle {
	return geometry.NewRectangleFromPoints(m.leftButtonDownPoint, m.position())
}

func (m *Mouse) position() geometry.Point {
	x, y := ebiten.CursorPosition()
	return geometry.NewPoint(float64(x), float64(y))
}

func (m *Mouse) getMouseDrawOptions(mouseRect geometry.Rectangle) *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(mouseRect.Point().X(), mouseRect.Point().Y())
	return opts
}

func (m *Mouse) getMouseImage(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if isCloseToEdge(i, width) || isCloseToEdge(j, height) {
				img.Set(i, j, color.White)
			}
		}
	}
	return img
}

func isCloseToEdge(i, j int) bool {
	return i == 0 || i == 1 || i == j-1 || i == j-2
}
