package core

import (
	"bufio"
	"fmt"
	"go_rts/client/render"
	"go_rts/types/geometry"
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
	RenderComponent   render.IRenderComponent
	PositionComponent geometry.IPositionComponent
	IsPathable        bool
}

// NewMapFromFile loads a map from a file
func NewMapFromFile(ssl render.ISpriteSheetLibrary, camera render.ICamera, filePath string) []*Tile {
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
			tiles = append(tiles, convertCharacterToTile(char, ssl, camera, point))
		}
	}

	return tiles
}

func convertCharacterToTile(char string, ssl render.ISpriteSheetLibrary, camera render.ICamera, point geometry.Point) *Tile {
	if char == "A" {
		return NewGrassTile(ssl, camera, point)
	} else if char == "B" {
		return NewWaterTile(ssl, camera, point)
	}
	return nil
}

// NewTile is a shortcut for creating a Tile
func NewTile(ssl render.ISpriteSheetLibrary, cam render.ICamera, name string, isPathable bool, p geometry.Point) *Tile {
	return &Tile{
		RenderComponent:   render.NewRenderComponent(ssl, cam, name),
		PositionComponent: geometry.NewPositionComponent(geometry.NewRectangle(tileWidth, tileHeight, p.X, p.Y), 0.0),
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

// ShrinkMapRectangle shrinks the mapRect by the percentage
func ShrinkMapRectangle(mapRect geometry.Rectangle, percentage float64) geometry.Rectangle {
	if percentage > 1 {
		panic(fmt.Sprintf("shrink percentage %v should be less than 1", percentage))
	}

	newWidth := mapRect.Width * percentage
	newHeight := mapRect.Height * percentage

	return geometry.NewRectangle(
		newWidth,
		newHeight,
		mapRect.Point.X+(mapRect.Width-newWidth)/2,
		mapRect.Point.Y+(mapRect.Height-newHeight)/2,
	)
}
