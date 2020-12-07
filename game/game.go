package game

import (
	"fmt"
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game an implementation of the interface provided by the ebiten v2 library
type Game struct {
	container     *Container
	gameMap       *GameMap
	units         []*Unit
	selectedUnits []*Unit

	leftButtonPressedListener  chan geometry.Point
	leftButtonReleasedListener chan geometry.Rectangle
}

// NewGame a shorcut method to instantiate a game object
func NewGame() Game {
	c := &Container{}
	game := Game{
		container:                  c,
		gameMap:                    NewMap(c.GetSpriteSheetLibrary(), c.GetCamera()),
		units:                      []*Unit{NewUnit(c.GetSpriteSheetLibrary(), c.GetCamera())},
		selectedUnits:              []*Unit{},
		leftButtonPressedListener:  make(chan geometry.Point, 1),
		leftButtonReleasedListener: make(chan geometry.Rectangle, 1),
	}

	game.container.GetMouse().LeftButtonPressedEvent().Subscribe(game.leftButtonPressedListener)
	game.container.GetMouse().LeftButtonReleasedEvent().Subscribe(game.leftButtonReleasedListener)

	return game
}

// Update is used to update all the game logic
func (g *Game) Update() error {
	g.updateCameraPosition()
	g.container.GetMouse().Update()
	g.listenForEvents()
	for _, u := range g.units {
		u.PositionComponent.MoveTowardsDestination()
	}
	return nil
}

func (g *Game) listenForEvents() {
	select {
	case rect := <-g.leftButtonReleasedListener:
		g.selectedUnits = selectUnits(rect, g.container.GetCamera(), g.units)
		fmt.Printf("number of selected Units=%v\n", len(g.selectedUnits))
	case point := <-g.leftButtonPressedListener:
		for _, u := range g.selectedUnits {
			u.PositionComponent.SetDestination(point)
		}
		fmt.Printf("thing pressed here: %v\n", point)
	default:
	}
}

func selectUnits(selectionRect geometry.Rectangle, camera render.ICamera, units []*Unit) []*Unit {
	selectedUnits := make([]*Unit, 0)
	for _, unit := range units {
		cameraTranslation := camera.Translation()
		unitIsoRect := unit.GetDrawRectangle()
		unitIsoRect.Point.Translate(cameraTranslation.Inverse())
		if selectionRect.Intersects(unitIsoRect) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
}

func (g *Game) updateCameraPosition() {
	moves := g.container.GetCamera().GetCameraMovements()
	for _, move := range moves {
		g.container.GetCamera().Translation().Translate(move)
	}

	mapPoints := g.gameMap.getMapPoints(1.0)
	if !g.doesScreenContainPoints(mapPoints...) {
		for _, move := range moves {
			g.container.GetCamera().Translation().Translate(move.Inverse())
		}
	}
}

func (g *Game) doesScreenContainPoints(points ...geometry.Point) bool {
	width, height := ebiten.WindowSize()
	screenOrigin := g.container.GetCamera().Translation()
	screenRect := geometry.NewRectangle(float64(width), float64(height), screenOrigin.X, screenOrigin.Y)

	for _, p := range points {
		if screenRect.Contains(p) {
			return true
		}
	}
	return false
}

// Draw is used to draw any relevant images on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.Draw(screen)
	for _, unit := range g.units {
		unit.Draw(screen)
	}
	g.container.GetMouse().Draw(screen)
}

// Layout returns the layout of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
