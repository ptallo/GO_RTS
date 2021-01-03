package core

import (
	"go_rts/types/geometry"

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

	game.container.GetEventHandler().OnLBP(game.onLeftButtonPressed)
	game.container.GetEventHandler().OnLBR(game.onLeftButtonReleased)
	game.container.GetEventHandler().OnRBP(game.onRightButtonPressed)

	return game
}

func (g *Game) onLeftButtonPressed(p geometry.Point) {
	g.gameObjects.SelectedUnits = []*Unit{}
}

func (g *Game) onLeftButtonReleased(r geometry.Rectangle) {
	cameraTranslation := *g.container.GetCamera().Translation()
	r.Point.Translate(cameraTranslation)
	g.gameObjects.SelectedUnits = g.gameObjects.SelectUnits(r)
}

func (g *Game) onRightButtonPressed(p geometry.Point) {
	cameraTranslation := *g.container.GetCamera().Translation()
	p.Translate(cameraTranslation)
	g.gameObjects.SetUnitsDestinations(&p)
}

// Update is used to update all the game logic
func (g *Game) Update() error {
	g.container.GetCamera().UpdateCameraPosition(
		float64(g.width),
		float64(g.height),
		ShrinkMapRectangle(GetMapRectangle(g.gameObjects.Tiles), 4),
	)
	g.container.GetMouse().Update()
	g.container.GetEventHandler().Listen()
	g.gameObjects.Update()
	return nil
}

// Draw is used to draw any relevant images on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.gameObjects.DrawGameObjects(screen)
	g.container.GetMouse().Draw(screen)
}

// Layout returns the layout of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
