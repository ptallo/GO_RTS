package input

import (
	"go_rts/types/geometry"

	"github.com/ethereum/go-ethereum/event"
	"github.com/hajimehoshi/ebiten/v2"
)

// IEventHandler is responsible for handling all input events and responding to them
type IEventHandler interface {
	Listen()
	OnLBP(fn func(geometry.Point))
	OnLBR(fn func(geometry.Rectangle))
	OnRBP(fn func(geometry.Point))
}

// IMouse defines an interface for wrapping any mouse system
type IMouse interface {
	Update()
	Draw(*ebiten.Image)
	LeftButtonPressedEvent() *event.Feed
	LeftButtonReleasedEvent() *event.Feed
	RightButtonPressedEvent() *event.Feed
}
