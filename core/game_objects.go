package core

import (
	"encoding/json"
	"go_rts/core/geometry"
	"go_rts/core/objects"
	"go_rts/core/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameObjects is a contianer for holding game information
type GameObjects struct {
	Tiles         []*objects.Tile
	Units         []*objects.Unit
	SelectedUnits []*objects.Unit
}

// NewGameObjects returns a standard initialization of the gameobjects
func NewGameObjects() *GameObjects {
	tile1 := objects.NewTile("water", false, geometry.NewPoint(10.0, 10.0))
	tile2 := objects.NewTile("grass", true, geometry.NewPoint(10.0, 10.0))
	unit1 := objects.NewUnit(geometry.NewPoint(100.0, 200.0))
	unit2 := objects.NewUnit(geometry.NewPoint(100.0, 300.0))
	return &GameObjects{
		Tiles:         []*objects.Tile{tile1, tile2},
		Units:         []*objects.Unit{unit1, unit2},
		SelectedUnits: []*objects.Unit{},
	}
}

// Serialize turns the GameObjects object to a byte array
func (g *GameObjects) Serialize() []byte {
	bytes, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Deserialize turns a byte array into a GameObjects object
func (g *GameObjects) Deserialize(bytes []byte) *GameObjects {
	var gameObjects GameObjects
	err := json.Unmarshal(bytes, &gameObjects)
	if err != nil {
		panic(err)
	}
	return &gameObjects
}

// Equals checks for equality between these gameObjects and others
func (g *GameObjects) Equals(g2 *GameObjects) bool {
	return areAllTilesInOtherArray(g.Tiles, g2.Tiles) &&
		areAllTilesInOtherArray(g2.Tiles, g.Tiles) &&
		areAllUnitsInOtherArray(g.Units, g2.Units) &&
		areAllUnitsInOtherArray(g2.Units, g.Units) &&
		areAllUnitsInOtherArray(g.SelectedUnits, g2.SelectedUnits) &&
		areAllUnitsInOtherArray(g2.SelectedUnits, g.SelectedUnits)
}

func areAllTilesInOtherArray(arr1, arr2 []*objects.Tile) bool {
	for _, t1 := range arr1 {
		inOtherArray := false
		for _, t2 := range arr2 {
			if t1.Equals(*t2) {
				inOtherArray = true
			}
		}

		if !inOtherArray {
			return false
		}
	}
	return true
}

func areAllUnitsInOtherArray(arr1, arr2 []*objects.Unit) bool {
	for _, u1 := range arr1 {
		inOtherArray := false
		for _, u2 := range arr2 {
			if u1.Equals(*u2) {
				inOtherArray = true
			}
		}

		if !inOtherArray {
			return false
		}
	}
	return true
}

// DrawGameObjects is responsible for drawing the images on the screen
func (g *GameObjects) DrawGameObjects(screen *ebiten.Image, camera render.ICamera) {
	for _, tile := range g.Tiles {
		camera.Draw(screen, tile.RenderComponent, *tile.PositionComponent.Rectangle.Point)
	}

	for _, unit := range g.Units {
		camera.Draw(screen, unit.RenderComponent, *unit.PositionComponent.Rectangle.Point)
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

func (g *GameObjects) getCollidableComponents() []geometry.PositionComponent {
	nonPathableTiles := make([]geometry.PositionComponent, 0)
	for _, tile := range g.Tiles {
		if !tile.IsPathable {
			nonPathableTiles = append(nonPathableTiles, *tile.PositionComponent)
		}
	}
	return nonPathableTiles
}

// SelectUnits selects all g.Units which intersect with the selectionRect
func (g *GameObjects) SelectUnits(selectionRect geometry.Rectangle) []*objects.Unit {
	selectedUnits := make([]*objects.Unit, 0)
	for _, unit := range g.Units {
		if selectionRect.Intersects(*unit.PositionComponent.Rectangle) {
			selectedUnits = append(selectedUnits, unit)
		}
	}
	return selectedUnits
}
