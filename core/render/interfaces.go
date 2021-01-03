package render

import (
	"go_rts/core/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

// ICamera is the interface which is implemented to provide camera usage
type ICamera interface {
	Draw(*ebiten.Image, *RenderComponent, geometry.Point)
	UpdateCameraPosition(float64, float64, geometry.Rectangle)
	Translation() *geometry.Point
}

// ISpriteSheetLibrary is an interface which defines a central store for sprite sheets
type ISpriteSheetLibrary interface {
	GetSpriteSheet(string) *SpriteSheet
}
