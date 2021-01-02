package main

import (
	"fmt"
	"go_rts/types/geometry"
)

func main() {
	rect := geometry.NewRectangle(10.0, 10.0, 10.0, 10.0)
	fmt.Printf("%+v", rect)
}
