package main

import (
	"go_rts/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	width, height := 800, 600
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go RTS")
	if err := ebiten.RunGame(core.NewGame(width, height)); err != nil {
		panic(err)
	}
}
