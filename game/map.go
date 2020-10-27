package game

import (
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	tiles    []*Tile
	imageMap map[string]*ebiten.Image
}

type Tile struct {
	name  string
	point render.Point
}

func NewMap() *GameMap {
	tiles := make([]*Tile, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p := render.NewPoint(
				float64((i+1)*64),
				float64((j+1)*64),
			)

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

func (m *GameMap) DrawMap(camera *render.Camera, screen *ebiten.Image) {
	for i := range m.tiles {
		tile := m.tiles[i]
		imageToDraw := m.imageMap[tile.name]
		isoPoint := render.CartoToIso(tile.point)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(isoPoint.X(), isoPoint.Y())
		camera.DrawImage(screen, imageToDraw, opts)
	}
}

func NewTileNameToImageMap() map[string]*ebiten.Image {
	return map[string]*ebiten.Image{
		"grass": render.NewImageFromPath("./assets/tiles/grass.png"),
		"water": render.NewImageFromPath("./assets/tiles/water.png"),
	}
}
