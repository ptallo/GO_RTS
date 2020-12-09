package game

import (
	"go_rts/geometry"
)

// EventHandler is a container for all event listeners in the game
type EventHandler struct {
	LeftButtonPressedListener  chan geometry.Point
	LeftButtonReleasedListener chan geometry.Rectangle
	RightButtonPressedListener chan geometry.Point
}

// NewEventHandler initializes all relevent event listeners the game will need
func NewEventHandler() *EventHandler {
	return &EventHandler{
		LeftButtonPressedListener:  make(chan geometry.Point, 1),
		LeftButtonReleasedListener: make(chan geometry.Rectangle, 1),
		RightButtonPressedListener: make(chan geometry.Point, 1),
	}
}

// Subscribe makes sure all the event listeners are subscribed to their corresponding events
func (ev *EventHandler) Subscribe(mouse IMouse) {
	mouse.LeftButtonPressedEvent().Subscribe(ev.LeftButtonPressedListener)
	mouse.LeftButtonReleasedEvent().Subscribe(ev.LeftButtonReleasedListener)
	mouse.RightButtonPressedEvent().Subscribe(ev.RightButtonPressedListener)
}
