package game

import (
	"go_rts/components"
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game an implementation of the interface provided by the ebiten v2 library
type Game struct {
	container *Container

	tiles         []*Tile
	units         []*Unit
	selectedUnits []*Unit
}

// NewGame a shorcut method to instantiate a game object
func NewGame() Game {
	c := &Container{}
	game := Game{
		container: c,
		tiles:     NewMap(c.GetSpriteSheetLibrary(), c.GetCamera()),
		units: []*Unit{
			NewUnit(c.GetSpriteSheetLibrary(), c.GetCamera(), geometry.NewPoint(40.0, 40.0)),
			NewUnit(c.GetSpriteSheetLibrary(), c.GetCamera(), geometry.NewPoint(200.0, 200.0)),
		},
		selectedUnits: []*Unit{},
	}

	game.container.GetEventHandler().Subscribe(game.container.GetMouse())

	return game
}

// Update is used to update all the game logic
func (g *Game) Update() error {
	g.updateCameraPosition()
	g.container.GetMouse().Update()
	g.listenForEvents()

	for i, u := range g.units {
		otherUnits := make([]components.IPositionComponent, 0)
		for j := range g.units {
			if i != j {
				otherUnits = append(otherUnits, g.units[j].PositionComponent)
			}
		}
		u.PositionComponent.MoveTowardsDestination(otherUnits)
	}
	return nil
}

func (g *Game) listenForEvents() {
	select {
	case rect := <-g.container.GetEventHandler().LeftButtonReleasedListener:
		g.selectedUnits = selectUnits(rect, g.container.GetCamera(), g.units)
	case _ = <-g.container.GetEventHandler().LeftButtonPressedListener:
		g.selectedUnits = []*Unit{}
	default:
	}

	select {
	case point := <-g.container.GetEventHandler().RightButtonPressedListener:
		g.setUnitsDestination(&point)
	default:
	}
}

func selectUnits(selectionRect geometry.Rectangle, camera render.ICamera, units []*Unit) []*Unit {
	selectedUnits := make([]*Unit, 0)
	for _, unit := range units {
		cameraTranslation := camera.Translation()
		unitIsoRect := unit.RenderComponent.GetDrawRectangle(geometry.CartoToIso(*unit.PositionComponent.GetPosition()))
		unitIsoRect.Point.Translate(cameraTranslation.Inverse())
		if selectionRect.Intersects(unitIsoRect) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
}

func (g *Game) setUnitsDestination(p *geometry.Point) {
	tilePositionComponents := make([]components.IPositionComponent, 0)
	for _, tile := range g.tiles {
		tilePositionComponents = append(tilePositionComponents, tile.PositionComponent)
	}

	p.Translate(*g.container.GetCamera().Translation())
	destination := geometry.IsoToCarto(*p)
	for _, u := range g.selectedUnits {
		u.PositionComponent.SetDestination(destination, tilePositionComponents)
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
