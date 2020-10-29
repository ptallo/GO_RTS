package game

import (
	"go_rts/objects"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gameMap *GameMap
	camera  *render.Camera
}

func NewGame() Game {
	cam := render.NewCamera()
	return Game{
		gameMap: NewMap(),
		camera:  &cam,
	}
}

func (g *Game) Update() error {
	g.updateCameraPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.DrawMap(g.camera, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) updateCameraPosition() {
	width, height := ebiten.WindowSize()

	cursorX, cursorY := ebiten.CursorPosition()
	p := objects.NewPoint(float64(cursorX), float64(cursorY))

	moves := make([]objects.Point, 0)
	if p.X() < float64(width)*0.1 {
		moves = append(moves, objects.NewPoint(-g.camera.Speed(), 0))
	}
	if p.X() > float64(width)*0.9 {
		moves = append(moves, objects.NewPoint(g.camera.Speed(), 0))
	}
	if p.Y() < float64(height)*0.1 {
		moves = append(moves, objects.NewPoint(0, -g.camera.Speed()))
	}
	if p.Y() > float64(height)*0.9 {
		moves = append(moves, objects.NewPoint(0, g.camera.Speed()))
	}

	for _, move := range moves {
		g.camera.MoveCamera(move)
	}

	if !g.isScreenAndMapOverlapping() {
		for _, move := range moves {
			g.camera.MoveCamera(move.Inverse())
		}
	}
}

func (g *Game) isScreenAndMapOverlapping() bool {
	width, height := ebiten.WindowSize()
	screenOrigin := g.camera.Translation()
	screenRect := objects.NewRectangle(float64(width), float64(height), screenOrigin.X(), screenOrigin.Y())
	mapRect := g.gameMap.GetMapRectangle()
	return screenRect.IsOverlapping(mapRect)
}
