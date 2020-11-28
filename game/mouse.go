package game

import (
	"fmt"
	"go_rts/geometry"
	"go_rts/render"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mouse struct {
	camera                 *render.Camera
	leftButtonDownDuration int
	leftButtonDownPoint    geometry.Point
}

func NewMouse(camera *render.Camera) *Mouse {
	m := Mouse{
		camera:                 camera,
		leftButtonDownDuration: 0,
		leftButtonDownPoint:    geometry.NewPoint(0, 0),
	}
	return &m
}

func (m *Mouse) Update(units []*Unit) []*Unit {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		m.leftButtonDownDuration++
		if m.leftButtonDownDuration == 1 {
			m.leftButtonDownPoint = m.position()
		}
	} else {
		selectedUnits := make([]*Unit, 0)
		if m.leftButtonDownDuration != 0 {
			selectedUnits = m.selectUnits(units)
			fmt.Printf("selected units length: %v\n", len(selectedUnits))
		}
		m.leftButtonDownDuration = 0
		return selectedUnits
	}

	return []*Unit{}
}

func (m *Mouse) selectUnits(units []*Unit) []*Unit {
	selectedUnits := make([]*Unit, 0)
	for _, unit := range units {
		cameraTranslation := m.camera.Translation()
		unitIsoRect := unit.GetDrawRectangle()
		unitIsoRect.Translate(cameraTranslation.Inverse())
		if m.getMouseSelectionRect().Intersects(unitIsoRect) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
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
		img := getMouseImage(int(rect.Width()), int(rect.Height()))
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

func getMouseImage(width, height int) *ebiten.Image {
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
