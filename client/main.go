package main

import (
	"encoding/gob"
	"go_rts/core"
	"go_rts/core/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gob.Register(geometry.Point{})
	gob.Register(geometry.Rectangle{})

	width, height := 800, 600
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go RTS")
	if err := ebiten.RunGame(core.NewGame(width, height)); err != nil {
		panic(err)
	}
}
