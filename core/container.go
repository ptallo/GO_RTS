package core

import (
	"go_rts/core/geometry"
	"go_rts/core/input"
	"go_rts/core/networking"
	"go_rts/core/render"
)

// Container is an object which holds objects for the game
type Container struct {
	spriteSheetLibrary render.ISpriteSheetLibrary
	camera             render.ICamera
	mouse              input.IMouse
	eventHandler       *EventHandler
	tcpClient          *networking.TCPClient
}

// GetSpriteSheetLibrary will lazy-load a singleton SpriteSheetLibrary object
func (c *Container) GetSpriteSheetLibrary() render.ISpriteSheetLibrary {
	if c.spriteSheetLibrary == nil {
		man, _ := render.NewSpriteSheet("units", "man")
		woman, _ := render.NewSpriteSheet("units", "woman")
		grass, _ := render.NewSpriteSheet("tiles", "grass")
		water, _ := render.NewSpriteSheet("tiles", "water")
		mouse, _ := render.NewSpriteSheet("ui", "mouse")
		c.spriteSheetLibrary = &render.SpriteSheetLibrary{
			Library: map[string]*render.SpriteSheet{
				"man":   man,
				"woman": woman,
				"grass": grass,
				"water": water,
				"mouse": mouse,
			},
		}
	}
	return c.spriteSheetLibrary
}

// GetCamera will lazy-load a singleton camera object
func (c *Container) GetCamera() render.ICamera {
	if c.camera == nil {
		c.camera = render.NewCamera(
			c.GetSpriteSheetLibrary(),
			geometry.NewPoint(0, 0),
			5.0,
		)
	}
	return c.camera
}

// GetMouse will lazy-load a singleton mouse object
func (c *Container) GetMouse() input.IMouse {
	if c.mouse == nil {
		c.mouse = input.NewMouse()
	}
	return c.mouse
}

// GetTCPClient will lazy-load a TCPClient with a live tcp connection
func (c *Container) GetTCPClient() *networking.TCPClient {
	if c.tcpClient == nil {
		c.tcpClient = networking.NewTCPClient()
	}
	return c.tcpClient
}

// GetEventHandler will lazy-load a singleton EventHandler
func (c *Container) GetEventHandler() *EventHandler {
	if c.eventHandler == nil {
		c.eventHandler = NewEventHandler(c.GetMouse(), c.GetTCPClient())
	}
	return c.eventHandler
}
