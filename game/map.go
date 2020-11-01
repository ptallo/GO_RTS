package game

import (
	"go_rts/geometry"
	"go_rts/render"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	tiles      []*Tile
	imageMap   map[string]*ebiten.Image
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
	n := 20
	tileW := 64.0
	tileH := 64.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p := geometry.NewPoint(
				float64(i+1)*tileW,
				float64(j+1)*tileH,
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

	return &GameMap{
		tiles:      tiles,
		imageMap:   NewTileNameToImageMap(),
		tileNum:    n,
		tileWidth:  tileW,
		tileHeight: tileH,
	}
}

func (m *GameMap) DrawMap(camera *render.Camera, screen *ebiten.Image) {
	for i := range m.tiles {
		tile := m.tiles[i]
		imageToDraw := m.imageMap[tile.name]
		isoPoint := geometry.CartoToIso(tile.point)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())
		camera.DrawImage(screen, imageToDraw, opts)
	}
}

func (m *GameMap) GetMapRectangle() geometry.Rectangle {
	p1, p2, p3, p4 := m.getMapCorners()
	minX := math.Min(math.Min(p1.X(), p2.X()), math.Min(p3.X(), p4.X()))
	maxX := math.Max(math.Max(p1.X(), p2.X()), math.Max(p3.X(), p4.X()))
	minY := math.Min(math.Min(p1.Y(), p2.Y()), math.Min(p3.Y(), p4.Y()))
	maxY := math.Max(math.Max(p1.Y(), p2.Y()), math.Max(p3.Y(), p4.Y()))
	return geometry.NewRectangle(maxX-minX, maxY-minY, minX, minY)
}

func (m *GameMap) getMapCorners() (geometry.Point, geometry.Point, geometry.Point, geometry.Point) {
	n := float64(m.tileNum)
	tileIn := 2.0
	minX := m.tileWidth * tileIn
	minY := m.tileHeight * tileIn
	maxX := m.tileWidth * (n - tileIn)
	maxY := m.tileHeight * (n - tileIn)
	p1 := geometry.CartoToIso(geometry.NewPoint(minX, minY))
	p2 := geometry.CartoToIso(geometry.NewPoint(minX, maxY))
	p3 := geometry.CartoToIso(geometry.NewPoint(maxX, minY))
	p4 := geometry.CartoToIso(geometry.NewPoint(maxX, maxY))
	return p1, p2, p3, p4
}

func NewTileNameToImageMap() map[string]*ebiten.Image {
	return map[string]*ebiten.Image{
		"grass": render.NewImageFromPath("./assets/tiles/grass.png"),
		"water": render.NewImageFromPath("./assets/tiles/water.png"),
	}
}
