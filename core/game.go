package core

import (
	"go_rts/core/geometry"
	"go_rts/core/networking"
	"go_rts/core/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game an implementation of the interface provided by the ebiten v2 library
type Game struct {
	width       int
	height      int
	container   *Container
	gameObjects *objects.GameObjects
}

// NewGame a shorcut method to instantiate a game object
func NewGame(width, height int) *Game {
	game := &Game{
		width:       width,
		height:      height,
		container:   &Container{},
		gameObjects: &objects.GameObjects{},
	}

	game.RegisterEvents()
	go game.container.GetTCPClient().ListenForGameObjects()

	return game
}

// RegisterEvents registers the corresponding events from the event handler with the game
func (g *Game) RegisterEvents() {
	g.container.GetEventHandler().OnLBP(func(p geometry.Point) {
		comm := networking.NewDeselectUnitsCommand(p)
		g.container.GetTCPClient().SendCommand(comm)
	})

	g.container.GetEventHandler().OnLBR(func(r geometry.Rectangle) {
		r = r.Move(g.container.GetCamera().Translation())
		comm := networking.NewSelectUnitsCommand(r)
		g.container.GetTCPClient().SendCommand(comm)
	})

	g.container.GetEventHandler().OnRBP(func(p geometry.Point) {
		p = p.Move(g.container.GetCamera().Translation())
		comm := networking.NewSetDestinationCommand(p)
		g.container.GetTCPClient().SendCommand(comm)
	})

	g.container.GetEventHandler().OnGameObjectsChanged(func(gameObjs objects.GameObjects) {
		g.gameObjects = &gameObjs
	})
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
