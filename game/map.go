package game

import (
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	tiles    []*Tile
	imageMap map[string]*ebiten.Image
}

func NewMap() *GameMap {
	tiles := make([]*Tile, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p := Point{
				x: float64((i + 1) * 64),
				y: float64((j + 1) * 64),
			}
			p = CartoToIso(p)

			tile := Tile{
				name:  "grass",
				point: p,
			}

			tiles = append(tiles, &tile)
		}
	}

	return &GameMap{
		tiles:    tiles,
		imageMap: NewTileNameToImageMap(),
	}
}

func (m *GameMap) DrawMap(screen *ebiten.Image) {
	for i := range m.tiles {
		tile := m.tiles[i]
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(tile.point.x, tile.point.y)
		screen.DrawImage(m.imageMap[tile.name], opts)
	}
}

type Tile struct {
	name  string
	point Point
}

func NewTileNameToImageMap() map[string]*ebiten.Image {
	return map[string]*ebiten.Image{
		"grass": render.NewImageFromPath("./assets/tiles/grass.png"),
		"water": render.NewImageFromPath("./assets/tiles/water.png"),
	}
}
