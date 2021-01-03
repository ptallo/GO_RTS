package core

import (
	"go_rts/types/geometry"
)

// EventHandler is a container for all event listeners in the game
type EventHandler struct {
	mouse IMouse

	LeftButtonPressedListener  chan geometry.Point
	leftButtonPressedFunctions []func(geometry.Point)

	LeftButtonReleasedListener  chan geometry.Rectangle
	leftButtonReleasedFunctions []func(geometry.Rectangle)

	RightButtonPressedListener  chan geometry.Point
	rightButtonPressedFunctions []func(geometry.Point)
}

// NewEventHandler initializes all relevent event listeners the game will need
func NewEventHandler(mouse IMouse) *EventHandler {
	ev := &EventHandler{
		LeftButtonPressedListener:  make(chan geometry.Point, 1),
		LeftButtonReleasedListener: make(chan geometry.Rectangle, 1),
		RightButtonPressedListener: make(chan geometry.Point, 1),
	}

	mouse.LeftButtonPressedEvent().Subscribe(ev.LeftButtonPressedListener)
	mouse.LeftButtonReleasedEvent().Subscribe(ev.LeftButtonReleasedListener)
	mouse.RightButtonPressedEvent().Subscribe(ev.RightButtonPressedListener)

	return ev
}

// Listen checks for all events and fires the related functions when they fire
func (ev *EventHandler) Listen() {
	select {
	case rect := <-ev.LeftButtonReleasedListener:
		for _, fn := range ev.leftButtonReleasedFunctions {
			go fn(rect)
		}
	default:
	}

	select {
	case point := <-ev.LeftButtonPressedListener:
		for _, fn := range ev.leftButtonPressedFunctions {
			go fn(point)
		}
	default:
	}

	select {
	case point := <-ev.RightButtonPressedListener:
		for _, fn := range ev.rightButtonPressedFunctions {
			go fn(point)
		}
	default:
	}
}

// OnLBP executes the passed in function when the LeftButtonPressed event fires
func (ev *EventHandler) OnLBP(fn func(geometry.Point)) {
	ev.leftButtonPressedFunctions = append(ev.leftButtonPressedFunctions, fn)
}

// OnLBR executes the passed in function when the LeftButtonReleased event fires
func (ev *EventHandler) OnLBR(fn func(geometry.Rectangle)) {
	ev.leftButtonReleasedFunctions = append(ev.leftButtonReleasedFunctions, fn)
}

// OnRBP executes the passed in function when the RightButtonPressed event fires
func (ev *EventHandler) OnRBP(fn func(geometry.Point)) {
	ev.rightButtonPressedFunctions = append(ev.rightButtonPressedFunctions, fn)
}
