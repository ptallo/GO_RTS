package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game an implementation of the interface provided by the ebiten v2 library
type Game struct {
	container    *Container
	eventHandler *EventHandler

	tiles         []*Tile
	units         []*Unit
	selectedUnits []*Unit
}

// NewGame a shorcut method to instantiate a game object
func NewGame() Game {
	c := &Container{}
	game := Game{
		container:     c,
		eventHandler:  NewEventHandler(),
		tiles:         NewMap(c.GetSpriteSheetLibrary(), c.GetCamera()),
		units:         []*Unit{NewUnit(c.GetSpriteSheetLibrary(), c.GetCamera())},
		selectedUnits: []*Unit{},
	}

	game.eventHandler.Subscribe(game.container.GetMouse())

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
	case rect := <-g.eventHandler.LeftButtonReleasedListener:
		g.selectedUnits = selectUnits(rect, g.container.GetCamera(), g.units)
	case _ = <-g.eventHandler.LeftButtonPressedListener:
		g.selectedUnits = []*Unit{}
	default:
	}

	select {
	case point := <-g.eventHandler.RightButtonPressedListener:
		g.setUnitsDestination(&point)
	default:
	}
}

func selectUnits(selectionRect geometry.Rectangle, camera render.ICamera, units []*Unit) []*Unit {
	selectedUnits := make([]*Unit, 0)
	for _, unit := range units {
		cameraTranslation := camera.Translation()
		unitIsoRect := unit.RenderComponent.GetDrawRectangle(*unit.PositionComponent.GetPosition())
		unitIsoRect.Point.Translate(cameraTranslation.Inverse())
		if selectionRect.Intersects(unitIsoRect) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
}

func (g *Game) setUnitsDestination(p *geometry.Point) {
	p.Translate(*g.container.GetCamera().Translation())
	destination := geometry.IsoToCarto(*p)
	for _, u := range g.selectedUnits {
		u.PositionComponent.SetDestination(destination)
	}
}

func (g *Game) updateCameraPosition() {
	moves := g.container.GetCamera().GetCameraMovements()
	for _, move := range moves {
		g.container.GetCamera().Translation().Translate(move)
	}

	mapPoints := make([]geometry.Point, 0)
	for _, tile := range g.tiles {
		mapPoints = append(mapPoints, tile.GetIsometricTileCorners()...)
	}

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
	for _, tile := range g.tiles {
		tile.RenderComponent.Draw(screen, geometry.CartoToIso(*tile.PositionComponent.GetPosition()))
	}

	for _, unit := range g.units {
		unit.RenderComponent.Draw(screen, geometry.CartoToIso(*unit.PositionComponent.GetPosition()))
	}
	g.container.GetMouse().Draw(screen)
}

// Layout returns the layout of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
