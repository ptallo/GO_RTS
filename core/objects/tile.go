package objects

import (
	"bufio"
	"go_rts/core/geometry"
	"go_rts/core/render"
	"os"
	"strings"
)

const (
	tileNum    = 10
	tileWidth  = 64.0
	tileHeight = 64.0
)

// Tile is an object describing a map tile
type Tile struct {
	RenderComponent   *render.RenderComponent
	PositionComponent geometry.IPositionComponent
	IsPathable        bool
}

// Equals checks whether this tile is equal to another tile
func (t Tile) Equals(t2 Tile) bool {
	return t.RenderComponent.Equals(*t2.RenderComponent) && t.PositionComponent.Equals(t2.PositionComponent) && t.IsPathable == t2.IsPathable
}

// NewMapFromFile loads a map from a file
func NewMapFromFile(filePath string) []*Tile {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tiles := make([]*Tile, 0)
	i := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		for j, char := range strings.Split(line, "") {
			point := geometry.NewPoint(float64(i)*tileWidth, float64(j)*tileHeight)
			tiles = append(tiles, convertCharacterToTile(char, point))
		}
	}

	return tiles
}

func convertCharacterToTile(char string, point geometry.Point) *Tile {
	if char == "A" {
		return newGrassTile(point)
	} else if char == "B" {
		return newWaterTile(point)
	}
	return nil
}

// NewTile is a shortcut for creating a Tile
func NewTile(name string, isPathable bool, p geometry.Point) *Tile {
	return &Tile{
		RenderComponent:   render.NewRenderComponent(name),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(tileWidth, tileHeight, p.X, p.Y), 0.0),
		IsPathable:        isPathable,
	}
}

func newWaterTile(p geometry.Point) *Tile {
	return NewTile("water", false, p)
}

// newGrassTile is a shortcut to create a grass tile
func newGrassTile(p geometry.Point) *Tile {
	return NewTile("grass", true, p)
}

// GetMapRectangle returns the rectangle describing the map
func GetMapRectangle(tiles []*Tile) geometry.Rectangle {
	minX := 999999999.0
	maxX := -999999999.0
	minY := 999999999.0
	maxY := -999999999.0

	for _, tile := range tiles {
		tileRect := tile.PositionComponent.GetRectangle()
		if minX > tileRect.Point.X {
			minX = tileRect.Point.X
		} else if maxX < tileRect.Point.X+tileRect.Width {
			maxX = tileRect.Point.X + tileRect.Width
		}

		if minY > tileRect.Point.Y {
			minY = tileRect.Point.Y
		} else if maxY < tileRect.Point.Y+tileRect.Height {
			maxY = tileRect.Point.Y + tileRect.Height
		}
	}

	return geometry.NewRectangleFromPoints(
		geometry.NewPoint(minX, minY),
		geometry.NewPoint(maxX, maxY),
	)
}

// ShrinkMapRectangle shrinks the mapRect by the number of tiles given
func ShrinkMapRectangle(mapRect geometry.Rectangle, numTiles int) geometry.Rectangle {
	numTilesF := float64(numTiles)
	return geometry.NewRectangle(
		mapRect.Width-(tileWidth*numTilesF),
		mapRect.Height-(tileHeight*numTilesF),
		mapRect.Point.X+(tileWidth*numTilesF/2),
		mapRect.Point.Y+(tileHeight*numTilesF/2),
	)
}
