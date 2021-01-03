package input

import (
	"github.com/ethereum/go-ethereum/event"
	"github.com/hajimehoshi/ebiten/v2"
)

// IMouse defines an interface for wrapping any mouse system
type IMouse interface {
	Update()
	Draw(*ebiten.Image)
	LeftButtonPressedEvent() *event.Feed
	LeftButtonReleasedEvent() *event.Feed
	RightButtonPressedEvent() *event.Feed
}
