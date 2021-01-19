package main

import (
	"encoding/gob"
	"fmt"
	"go_rts/core/geometry"
	"go_rts/core/networking"
	"go_rts/core/objects"
	"time"
)

func main() {
	gob.Register(geometry.Point{})
	gob.Register(geometry.Rectangle{})

	s := networking.NewTCPServer()

	selectUnitsListener := make(chan geometry.Rectangle, 1)
	s.SelectUnitsEvent.Subscribe(selectUnitsListener)

	deselectUnitsListener := make(chan geometry.Point, 1)
	s.DeselectUnitsEvent.Subscribe(deselectUnitsListener)

	setDestinationListener := make(chan geometry.Point, 1)
	s.SetDestinationEvent.Subscribe(setDestinationListener)

	go s.ListenForConnections()

	gameObjects := objects.NewGameObjects()

	for {
		time.Sleep(time.Second / 60)

		select {
		case r := <-selectUnitsListener:
			fmt.Printf("selecting units: %+v\n", r)
			gameObjects.SelectedUnits = gameObjects.SelectUnits(r)
			fmt.Printf("selected units: %+v\n", gameObjects.SelectedUnits)
		default:
		}

		select {
		case _ = <-deselectUnitsListener:
			fmt.Printf("deselecting units\n")
			gameObjects.SelectedUnits = []int{}
		default:
		}

		select {
		case p := <-setDestinationListener:
			fmt.Printf("Setting Destination %+v for units %+v \n", p, gameObjects.SelectedUnits)
			gameObjects.SetUnitsDestinations(&p)
		default:
		}

		gameObjects.Update()
		s.SendGameObjects(*gameObjects)
	}
}
