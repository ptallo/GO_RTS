package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Go RTS")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
