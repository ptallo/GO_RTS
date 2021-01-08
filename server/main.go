package main

import (
	"go_rts/core/geometry"
	"go_rts/core/networking"
	"go_rts/core/objects"
	"time"
)

func main() {
	s := networking.NewTCPServer()
	go s.Listen()

	gameObjects := objects.NewGameObjects()
	for _, u := range gameObjects.Units {
		u.PositionComponent.SetDestination(geometry.NewPoint(500.0, 300.0), objects.GetMapRectangle(gameObjects.Tiles), gameObjects.GetCollidableComponents())
	}

	for {
		time.Sleep(time.Second / 60)
		gameObjects.Update()
		s.SendGameObjects(*gameObjects)
	}
}
