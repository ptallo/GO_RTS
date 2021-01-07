package core

import (
	"go_rts/core/geometry"
	"go_rts/core/input"
	"go_rts/core/networking"
	"go_rts/core/objects"
)

// IEventHandler is responsible for handling all input events and responding to them
type IEventHandler interface {
	Listen()
	OnLBP(fn func(geometry.Point))
	OnLBR(fn func(geometry.Rectangle))
	OnRBP(fn func(geometry.Point))
}

// EventHandler is a container for all event listeners in the game
type EventHandler struct {
	LeftButtonPressedListener  chan geometry.Point
	leftButtonPressedFunctions []func(geometry.Point)

	LeftButtonReleasedListener  chan geometry.Rectangle
	leftButtonReleasedFunctions []func(geometry.Rectangle)

	RightButtonPressedListener  chan geometry.Point
	rightButtonPressedFunctions []func(geometry.Point)

	GameObjectsChangedListener  chan *objects.GameObjects
	gameObjectsChangedFunctions []func(*objects.GameObjects)
}

// NewEventHandler initializes all relevent event listeners the game will need
func NewEventHandler(mouse input.IMouse, tcpClient *networking.TCPClient) *EventHandler {
	ev := &EventHandler{
		LeftButtonPressedListener:   make(chan geometry.Point, 1),
		leftButtonPressedFunctions:  []func(geometry.Point){},
		LeftButtonReleasedListener:  make(chan geometry.Rectangle, 1),
		leftButtonReleasedFunctions: []func(geometry.Rectangle){},
		RightButtonPressedListener:  make(chan geometry.Point, 1),
		rightButtonPressedFunctions: []func(geometry.Point){},
		GameObjectsChangedListener:  make(chan *objects.GameObjects, 1),
		gameObjectsChangedFunctions: []func(*objects.GameObjects){},
	}

	mouse.LeftButtonPressedEvent().Subscribe(ev.LeftButtonPressedListener)
	mouse.LeftButtonReleasedEvent().Subscribe(ev.LeftButtonReleasedListener)
	mouse.RightButtonPressedEvent().Subscribe(ev.RightButtonPressedListener)
	tcpClient.GameObjectsChangedEvent().Subscribe(ev.GameObjectsChangedListener)

	return ev
}

// Listen checks for all events and fires the related functions when they fire
func (ev *EventHandler) Listen() {
	select {
	case rect := <-ev.LeftButtonReleasedListener:
		for _, fn := range ev.leftButtonReleasedFunctions {
			fn(rect)
		}
	default:
	}

	select {
	case point := <-ev.LeftButtonPressedListener:
		for _, fn := range ev.leftButtonPressedFunctions {
			fn(point)
		}
	default:
	}

	select {
	case point := <-ev.RightButtonPressedListener:
		for _, fn := range ev.rightButtonPressedFunctions {
			fn(point)
		}
	default:
	}

	select {
	case gameObjects := <-ev.GameObjectsChangedListener:
		for _, fn := range ev.gameObjectsChangedFunctions {
			fn(gameObjects)
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

// OnGameObjectsChanged subscribes the 'fn' argument to be called with the GameObjectsChanged event fires
func (ev *EventHandler) OnGameObjectsChanged(fn func(*objects.GameObjects)) {
	ev.gameObjectsChangedFunctions = append(ev.gameObjectsChangedFunctions, fn)
}
