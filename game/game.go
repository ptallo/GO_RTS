package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	camera         *render.Camera
	spriteSheetLib map[string]*render.SpriteSheet
	mouse          *Mouse
	gameMap        *GameMap
	unit           *Unit
}

func NewGame() Game {
	ssl := render.NewSpriteSheetMap()
	return Game{
		camera:         render.NewCamera(),
		spriteSheetLib: ssl,
		mouse:          NewMouse(ssl["mouse"]),
		gameMap:        NewMap(),
		unit:           NewUnit(),
	}
}

func (g *Game) Update() error {
	g.updateCameraPosition()
	g.mouse.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.Draw(g.camera, screen, g.spriteSheetLib)
	g.unit.Draw(g.camera, screen, g.spriteSheetLib)
	g.mouse.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) updateCameraPosition() {
	moves := g.camera.GetCameraMovements()
	for _, move := range moves {
		g.camera.MoveCamera(move)
	}

	mapPoints := g.gameMap.getMapPoints(1.0)
	if !g.doesScreenContainAPoint(mapPoints...) {
		for _, move := range moves {
			g.camera.MoveCamera(move.Inverse())
		}
	}
}

func (g *Game) doesScreenContainAPoint(points ...geometry.Point) bool {
	width, height := ebiten.WindowSize()
	screenOrigin := g.camera.Translation()
	screenRect := geometry.NewRectangle(float64(width), float64(height), screenOrigin.X(), screenOrigin.Y())

	for _, p := range points {
		if screenRect.Contains(p) {
			return true
		}
	}
	return false
}
