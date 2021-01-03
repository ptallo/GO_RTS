package core

import (
	"go_rts/core/geometry"
	"go_rts/core/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game an implementation of the interface provided by the ebiten v2 library
type Game struct {
	width       int
	height      int
	container   *Container
	gameObjects *GameObjects
}

// NewGame a shorcut method to instantiate a game object
func NewGame(width, height int) Game {
	c := &Container{}
	game := Game{
		width:       width,
		height:      height,
		container:   c,
		gameObjects: NewGameObjects(c),
	}

	game.container.GetEventHandler().OnLBP(func(p geometry.Point) {
		game.gameObjects.SelectedUnits = []*objects.Unit{}
	})

	game.container.GetEventHandler().OnLBR(func(r geometry.Rectangle) {
		cameraTranslation := *game.container.GetCamera().Translation()
		r.Point.Translate(cameraTranslation)
		game.gameObjects.SelectedUnits = game.gameObjects.SelectUnits(r)
	})

	game.container.GetEventHandler().OnRBP(func(p geometry.Point) {
		cameraTranslation := *game.container.GetCamera().Translation()
		p.Translate(cameraTranslation)
		game.gameObjects.SetUnitsDestinations(&p)
	})

	return game
}

// Update is used to update all the game logic
func (g *Game) Update() error {
	g.container.GetCamera().UpdateCameraPosition(
		float64(g.width),
		float64(g.height),
		objects.ShrinkMapRectangle(objects.GetMapRectangle(g.gameObjects.Tiles), 4),
	)
	g.container.GetMouse().Update()
	g.container.GetEventHandler().Listen()
	g.gameObjects.Update()
	return nil
}

// Draw is used to draw any relevant images on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.gameObjects.DrawGameObjects(screen, g.container.GetCamera())
	g.container.GetMouse().Draw(screen)
}

// Layout returns the layout of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
