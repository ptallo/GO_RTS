package game

import (
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tileImage *ebiten.Image
}

func NewGame() Game {
	return Game{
		tileImage: render.NewImageFromPath("./assets/block.png"),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10.0, 10.0)

	screen.DrawImage(g.tileImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
