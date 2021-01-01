package main

import (
	"go_rts/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Go RTS")
	g := core.NewGame()
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
