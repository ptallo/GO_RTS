package game

import (
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gameMap *GameMap
	camera  *render.Camera
}

func NewGame() Game {
	cam := render.NewCamera()
	return Game{
		gameMap: NewMap(),
		camera:  &cam,
	}
}

func (g *Game) Update() error {
	g.camera.UpdateCameraPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.DrawMap(g.camera, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
