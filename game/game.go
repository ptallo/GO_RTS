package game

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	camera         *render.Camera
	spriteSheetLib map[string]*render.SpriteSheet
	gameMap        *GameMap
	unit           *Unit
}

func NewGame() Game {
	cam := render.NewCamera()
	gameMap := NewMap()
	unit := NewUnit()

	return Game{
		camera:         &cam,
		spriteSheetLib: render.NewSpriteSheetMap(),
		gameMap:        &gameMap,
		unit:           &unit,
	}
}

func (g *Game) Update() error {
	g.updateCameraPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.Draw(g.camera, screen, g.spriteSheetLib)
	g.unit.Draw(g.camera, screen, g.spriteSheetLib)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) updateCameraPosition() {
	width, height := ebiten.WindowSize()

	cursorX, cursorY := ebiten.CursorPosition()
	p := geometry.NewPoint(float64(cursorX), float64(cursorY))

	moves := make([]geometry.Point, 0)
	if p.X() < float64(width)*0.1 {
		moves = append(moves, geometry.NewPoint(-g.camera.Speed(), 0))
	}
	if p.X() > float64(width)*0.9 {
		moves = append(moves, geometry.NewPoint(g.camera.Speed(), 0))
	}
	if p.Y() < float64(height)*0.1 {
		moves = append(moves, geometry.NewPoint(0, -g.camera.Speed()))
	}
	if p.Y() > float64(height)*0.9 {
		moves = append(moves, geometry.NewPoint(0, g.camera.Speed()))
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
	// Get the rectangle describing the screen
	width, height := ebiten.WindowSize()
	screenOrigin := g.camera.Translation()
	screenRect := geometry.NewRectangle(float64(width), float64(height), screenOrigin.X(), screenOrigin.Y())

	points := g.gameMap.getMapPoints(1.0)

	for _, p := range points {
		if screenRect.Contains(p) {
			return true
		}
	}
	return false
}
