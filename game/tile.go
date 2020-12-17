package game

import (
	"go_rts/components"
	"go_rts/geometry"
	"go_rts/render"
)

const (
	tileNum    = 10
	tileWidth  = 64.0
	tileHeight = 64.0
)

// Tile is an object describing a map tile
type Tile struct {
	RenderComponent   components.IRenderComponent
	PositionComponent components.IPositionComponent
	IsPathable        bool
}

// NewMap is a shorcut for defining a GameMap object
func NewMap(ssl render.ISpriteSheetLibrary, camera render.ICamera) []*Tile {
	tiles := make([]*Tile, 0)
	for i := 0; i < tileNum; i++ {
		for j := 0; j < tileNum; j++ {
			p := geometry.NewPoint(
				float64(i)*tileWidth,
				float64(j)*tileHeight,
			)

			if i == 0 || j == 0 || i == tileNum-1 || j == tileNum-1 {
				tiles = append(tiles, NewWaterTile(ssl, camera, p))
			} else {
				tiles = append(tiles, NewGrassTile(ssl, camera, p))
			}

		}
	}

	return tiles
}

// NewTile is a shortcut for creating a Tile
func NewTile(ssl render.ISpriteSheetLibrary, cam render.ICamera, name string, isPathable bool, p geometry.Point) *Tile {
	return &Tile{
		RenderComponent:   components.NewRenderComponent(ssl, cam, name),
		PositionComponent: components.NewPositionComponent(geometry.NewRectangle(tileWidth, tileHeight, p.X, p.Y), 0.0),
		IsPathable:        isPathable,
	}
}

// NewWaterTile is a shortcut to create a water tile
func NewWaterTile(ssl render.ISpriteSheetLibrary, cam render.ICamera, p geometry.Point) *Tile {
	return NewTile(ssl, cam, "water", false, p)
}

// NewGrassTile is a shortcut to create a grass tile
func NewGrassTile(ssl render.ISpriteSheetLibrary, cam render.ICamera, p geometry.Point) *Tile {
	return NewTile(ssl, cam, "grass", true, p)
}

// GetIsometricTileCorners returns the four corners of the tile in isometric space
func (t *Tile) GetIsometricTileCorners() []geometry.Point {
	tileOrigin := t.PositionComponent.GetPosition()
	points := []geometry.Point{
		geometry.NewPoint(tileOrigin.X, tileOrigin.Y),
		geometry.NewPoint(tileOrigin.X+tileWidth, tileOrigin.Y),
		geometry.NewPoint(tileOrigin.X, tileOrigin.Y+tileHeight),
		geometry.NewPoint(tileOrigin.X+tileWidth, tileOrigin.Y+tileHeight),
	}
	return points
}
