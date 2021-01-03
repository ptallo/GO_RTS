package main

import (
	"go_rts/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// conn, err := net.Dial("tcp", "localhost:8080")
	// if err != nil {
	// 	fmt.Println("failed to connect to server")
	// }

	// for {
	// 	message, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	message = string(message)
	// 	fmt.Println(message)
	// 	if strings.Trim(message, "\r\n") == "DONE" {
	// 		break
	// 	}
	// }

	// conn.Close()

	width, height := 800, 600
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go RTS")
	g := core.NewGame(width, height)
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
