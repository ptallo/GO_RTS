package main

import (
	"fmt"
	"go_rts/types/core"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("failed to connect to server")
	}

	conn.Close()

	width, height := 800, 600

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go RTS")
	g := core.NewGame(width, height)
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
