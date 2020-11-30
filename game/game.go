package game

import (
	"go_rts/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	container *Container
	mouse     *Mouse
	gameMap   *GameMap
	units     []*Unit
}

func NewGame() Game {
	c := &Container{}
	return Game{
		container: c,
		mouse:     NewMouse(c.GetCamera()),
		gameMap:   NewMap(c.GetSpriteSheetLibrary(), c.GetCamera()),
		units:     []*Unit{NewUnit(c.GetSpriteSheetLibrary(), c.GetCamera())},
	}
}

func (g *Game) Update() error {
	g.updateCameraPosition()
	g.mouse.Update(g.units)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameMap.Draw(screen)
	for _, unit := range g.units {
		unit.Draw(screen)
	}
	g.mouse.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) updateCameraPosition() {
	moves := g.container.GetCamera().GetCameraMovements()
	for _, move := range moves {
		g.container.GetCamera().MoveCamera(move)
	}

	mapPoints := g.gameMap.getMapPoints(1.0)
	if !g.doesScreenContainAPoint(mapPoints...) {
		for _, move := range moves {
			g.container.GetCamera().MoveCamera(move.Inverse())
		}
	}
}

func (g *Game) doesScreenContainAPoint(points ...geometry.Point) bool {
	width, height := ebiten.WindowSize()
	screenOrigin := g.container.GetCamera().Translation
	screenRect := geometry.NewRectangle(float64(width), float64(height), screenOrigin.X, screenOrigin.Y)

	for _, p := range points {
		if screenRect.Contains(p) {
			return true
		}
	}
	return false
}
