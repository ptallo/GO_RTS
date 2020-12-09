package game

import (
	"go_rts/geometry"
	"go_rts/render"
)

// Container is an object which holds objects for the game
type Container struct {
	spriteSheetLibrary render.ISpriteSheetLibrary
	camera             render.ICamera
	mouse              IMouse
	eventHandler       *EventHandler
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
	p := geometry.NewPoint(0, 0)
	if c.camera == nil {
		c.camera = render.NewCamera(
			&p,
			5.0,
		)
	}
	return c.camera
}

// GetMouse will lazy-load a singleton mouse object
func (c *Container) GetMouse() IMouse {
	if c.mouse == nil {
		c.mouse = NewMouse()
	}
	return c.mouse
}

// GetEventHandler will lazy-load a singleton EventHandler
func (c *Container) GetEventHandler() *EventHandler {
	if c.eventHandler == nil {
		c.eventHandler = NewEventHandler()
	}
	return c.eventHandler
}
