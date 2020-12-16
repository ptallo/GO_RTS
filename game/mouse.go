package game

import (
	"go_rts/geometry"
	"image/color"

	"github.com/ethereum/go-ethereum/event"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:generate mockgen -destination=../mocks/mock_IMouse.go -package=mocks go_rts/game IMouse

// IMouse defines an interface for wrapping any mouse system
type IMouse interface {
	Update()
	Draw(*ebiten.Image)
	LeftButtonPressedEvent() *event.Feed
	LeftButtonReleasedEvent() *event.Feed
	RightButtonPressedEvent() *event.Feed
}

// Mouse is an object wrapping all ebiten mouse utilities
type Mouse struct {
	leftButtonDownDuration  int
	leftButtonDownPoint     geometry.Point
	leftButtonPressedEvent  *event.Feed
	leftButtonReleasedEvent *event.Feed
	rightButtonDownDuration int
	rightButtonDownPoint    geometry.Point
	rightButtonPressedEvent *event.Feed
}

// NewMouse is shorcut method to defining a Mouse object
func NewMouse() *Mouse {
	return &Mouse{
		leftButtonDownDuration:  0,
		leftButtonDownPoint:     geometry.NewPoint(0, 0),
		leftButtonPressedEvent:  &event.Feed{},
		leftButtonReleasedEvent: &event.Feed{},
		rightButtonDownDuration: 0,
		rightButtonDownPoint:    geometry.NewPoint(0, 0),
		rightButtonPressedEvent: &event.Feed{},
	}
}

// Update is responsible for firing events related to the mouse object
func (m *Mouse) Update() {
	m.fireEvents()
	m.updateCount()
}

func (m *Mouse) fireEvents() {
	if m.isLeftButtonJustPressed() {
		m.leftButtonDownPoint = m.getMousePosition()
		m.leftButtonPressedEvent.Send(m.leftButtonDownPoint)
	}

	if m.isLeftButtonJustReleased() {
		m.leftButtonReleasedEvent.Send(m.getMouseSelectionRect())
	}

	if m.isRightButtonJustPressed() {
		m.rightButtonDownPoint = m.getMousePosition()
		m.rightButtonPressedEvent.Send(m.rightButtonDownPoint)
	}
}

func (m *Mouse) isLeftButtonJustReleased() bool {
	return !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && m.leftButtonDownDuration != 0
}

func (m *Mouse) isLeftButtonJustPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && m.leftButtonDownDuration == 0
}

func (m *Mouse) isRightButtonJustReleased() bool {
	return !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && m.rightButtonDownDuration != 0
}

func (m *Mouse) isRightButtonJustPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && m.rightButtonDownDuration == 0
}

func (m *Mouse) updateCount() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		m.leftButtonDownDuration++
	} else {
		m.leftButtonDownDuration = 0
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		m.rightButtonDownDuration++
	} else {
		m.rightButtonDownDuration = 0
	}
}

// Draw is responsible for drawing any mouse related effects on the screen
func (m *Mouse) Draw(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		rect := m.getMouseSelectionRect()
		opts := m.getMouseDrawOptions(rect)
		img := getMouseImage(int(rect.Width), int(rect.Height))
		screen.DrawImage(img, opts)
	}
}

func (m *Mouse) getMouseSelectionRect() geometry.Rectangle {
	return geometry.NewRectangleFromPoints(m.leftButtonDownPoint, m.getMousePosition())
}

func (m *Mouse) getMousePosition() geometry.Point {
	x, y := ebiten.CursorPosition()
	return geometry.NewPoint(float64(x), float64(y))
}

func (m *Mouse) getMouseDrawOptions(mouseRect geometry.Rectangle) *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(mouseRect.Point.X, mouseRect.Point.Y)
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

// LeftButtonPressedEvent returns the event for pressing the left button
func (m *Mouse) LeftButtonPressedEvent() *event.Feed {
	return m.leftButtonPressedEvent
}

// LeftButtonReleasedEvent returns the event for releasing the left button
func (m *Mouse) LeftButtonReleasedEvent() *event.Feed {
	return m.leftButtonReleasedEvent
}

// RightButtonPressedEvent returns the event for pressing the right button
func (m *Mouse) RightButtonPressedEvent() *event.Feed {
	return m.rightButtonPressedEvent
}
