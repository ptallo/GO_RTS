package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	tiles      []*Tile
	tileNum    int
	tileWidth  float64
	tileHeight float64
}

type Tile struct {
	name  string
	point geometry.Point
}

func NewMap() *GameMap {
	tiles := make([]*Tile, 0)
	n := 10
	tileW := 64.0
	tileH := 64.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p := geometry.NewPoint(
				float64(i)*tileW,
				float64(j)*tileH,
			)

			tileName := "grass"
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				tileName = "water"
			}

			tile := Tile{
				name:  tileName,
				point: p,
			}

			tiles = append(tiles, &tile)
		}
	}

	gm := GameMap{
		tiles:      tiles,
		tileNum:    n,
		tileWidth:  tileW,
		tileHeight: tileH,
	}
	return &gm
}

func (m *GameMap) Draw(camera *render.Camera, screen *ebiten.Image, lib map[string]*render.SpriteSheet) {
	for _, tile := range m.tiles {
		isoPoint := geometry.CartoToIso(tile.point)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())
		spriteSheet := lib[tile.name]
		spriteSheet.Draw(screen, camera, opts)
	}
}

func (m *GameMap) getMapPoints(numTilesToGoIn float64) []geometry.Point {
	n := float64(m.tileNum)
	points := make([]geometry.Point, 0)
	for i := numTilesToGoIn; i < n-numTilesToGoIn; i++ {
		for j := numTilesToGoIn; j < n-numTilesToGoIn; j++ {
			ps := m.getTileCorners(i, j)
			points = append(points, ps...)
		}
	}
	return points
}

func (m *GameMap) getTileCorners(i, j float64) []geometry.Point {
	w := m.tileWidth
	h := m.tileHeight
	points := []geometry.Point{
		geometry.CartoToIso(geometry.NewPoint(i*w, j*h)),
		geometry.CartoToIso(geometry.NewPoint((i+1)*w, j*h)),
		geometry.CartoToIso(geometry.NewPoint(i*w, (j+1)*h)),
		geometry.CartoToIso(geometry.NewPoint((i+1)*w, (j+1)*h)),
	}
	return points
}
