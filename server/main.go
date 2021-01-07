package main

import (
	"fmt"
	"go_rts/core/geometry"
	"go_rts/core/objects"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("couldn't start server")
	}

	defer listener.Close()
	fmt.Println("listening on tcp port 8080...")

	gameObjects := objects.NewGameObjects()
	for _, u := range gameObjects.Units {
		u.PositionComponent.SetDestination(geometry.NewPoint(500.0, 300.0), objects.GetMapRectangle(gameObjects.Tiles), gameObjects.GetCollidableComponents())
	}

	gameObjects.Update()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("couldn't accept connection from %v\n", conn.LocalAddr().String())
		} else {
			fmt.Printf("connection accepted from %v\n", conn.LocalAddr().String())

			defer conn.Close()

			fmt.Fprintf(conn, fmt.Sprintf("%s\n", string(gameObjects.Serialize())))
		}
	}
}
