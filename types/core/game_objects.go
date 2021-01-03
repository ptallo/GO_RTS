package core

import (
	"go_rts/types/geometry"
	"go_rts/types/objects"
	"go_rts/types/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameObjects is a contianer for holding game information
type GameObjects struct {
	Tiles         []*objects.Tile
	Units         []*objects.Unit
	SelectedUnits []*objects.Unit
}

// NewGameObjects returns a standard initialization of the gameobjects
func NewGameObjects(container *Container) *GameObjects {
	return &GameObjects{
		Tiles: objects.NewMapFromFile(container.GetSpriteSheetLibrary(), container.GetCamera(), "./assets/maps/map1.txt"),
		Units: []*objects.Unit{
			objects.NewUnit(container.GetSpriteSheetLibrary(), container.GetCamera(), geometry.NewPoint(100.0, 200.0)),
			objects.NewUnit(container.GetSpriteSheetLibrary(), container.GetCamera(), geometry.NewPoint(100.0, 300.0)),
		},
		SelectedUnits: []*objects.Unit{},
	}
}

// DrawGameObjects is responsible for drawing the images on the screen
func (g *GameObjects) DrawGameObjects(screen *ebiten.Image, camera render.ICamera) {
	for _, tile := range g.Tiles {
		camera.Draw(screen, tile.RenderComponent, *tile.PositionComponent.GetRectangle().Point)
	}

	for _, unit := range g.Units {
		camera.Draw(screen, unit.RenderComponent, *unit.PositionComponent.GetRectangle().Point)
	}
}

// Update updates all relevent information for gameObjects
func (g *GameObjects) Update() {
	for _, u := range g.Units {
		u.PositionComponent.MoveTowardsDestination(g.getCollidableComponents())
	}
}

// SetUnitsDestinations sets the destination of all g.SelectedUnits to the point p
func (g *GameObjects) SetUnitsDestinations(p *geometry.Point) {
	mapRect := objects.GetMapRectangle(g.Tiles)
	collidables := g.getCollidableComponents()
	for _, u := range g.SelectedUnits {
		u.PositionComponent.SetDestination(*p, mapRect, collidables)
	}
}

func (g *GameObjects) getCollidableComponents() []geometry.IPositionComponent {
	nonPathableTiles := make([]geometry.IPositionComponent, 0)
	for _, tile := range g.Tiles {
		if !tile.IsPathable {
			nonPathableTiles = append(nonPathableTiles, tile.PositionComponent)
		}
	}
	return nonPathableTiles
}

// SelectUnits selects all g.Units which intersect with the selectionRect
func (g *GameObjects) SelectUnits(selectionRect geometry.Rectangle) []*objects.Unit {
	selectedUnits := make([]*objects.Unit, 0)
	for _, unit := range g.Units {
		if selectionRect.Intersects(unit.PositionComponent.GetRectangle()) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
}
