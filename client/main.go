package main

import (
	"go_rts/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	width, height := 800, 600
	g := core.NewGame(width, height)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go RTS")
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}

}
