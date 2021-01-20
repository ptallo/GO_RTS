package objects

import (
	"go_rts/core/geometry"
	"go_rts/core/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameObjects is a contianer for holding game information
type GameObjects struct {
	UIDHandler    *UIDHandler
	Tiles         []*Tile
	Units         []*Unit
	SelectedUnits []int
}

// NewGameObjects returns a standard initialization of the gameobjects
func NewGameObjects() *GameObjects {
	uidHandler := &UIDHandler{}
	unit1 := NewUnit(uidHandler.getNewUID(), geometry.NewPoint(100.0, 200.0))
	unit2 := NewUnit(uidHandler.getNewUID(), geometry.NewPoint(100.0, 300.0))
	return &GameObjects{
		Tiles:         NewMapFromFile("../assets/maps/map1.txt"),
		Units:         []*Unit{unit1, unit2},
		SelectedUnits: []int{},
	}
}

// DrawGameObjects is responsible for drawing the images on the screen
func (g *GameObjects) DrawGameObjects(screen *ebiten.Image, camera render.ICamera) {
	for _, tile := range g.Tiles {
		camera.Draw(screen, tile.RenderComponent, tile.PositionComponent.Rectangle.Point)
	}

	for _, unit := range g.Units {
		camera.Draw(screen, unit.RenderComponent, unit.PositionComponent.Rectangle.Point)
	}
}

// Update updates all relevent information for gameObjects
func (g *GameObjects) Update() {
	for _, u := range g.Units {
		u.PositionComponent.MoveTowardsDestination(g.GetCollidableComponents())
	}
}

// SetUnitsDestinations sets the destination of all g.SelectedUnits to the point p
func (g *GameObjects) SetUnitsDestinations(p *geometry.Point) {
	mapRect := GetMapRectangle(g.Tiles)
	collidables := g.GetCollidableComponents()
	for _, uid := range g.SelectedUnits {
		for _, unit := range g.Units {
			if unit.UID == uid {
				unit.PositionComponent.SetDestination(*p, mapRect, collidables)
			}
		}
	}
}

// GetCollidableComponents gets all the collidable components in the gameObjects
func (g *GameObjects) GetCollidableComponents() []geometry.PositionComponent {
	nonPathableTiles := make([]geometry.PositionComponent, 0)
	for _, tile := range g.Tiles {
		if !tile.IsPathable {
			nonPathableTiles = append(nonPathableTiles, *tile.PositionComponent)
		}
	}
	return nonPathableTiles
}

// SelectUnits selects all g.Units which intersect with the selectionRect
func (g *GameObjects) SelectUnits(selectionRect geometry.Rectangle) []int {
	selectedUnits := make([]int, 0)
	for _, unit := range g.Units {
		if selectionRect.Intersects(unit.PositionComponent.Rectangle) {
			selectedUnits = append(selectedUnits, unit.UID)
		}
	}
	return selectedUnits
}
