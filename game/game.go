package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gameMap *GameMap
}

func NewGame() Game {
	return Game{
		gameMap: NewMap(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.DrawMap(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
