package components

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

// IRenderComponent is an interface which describes drawing something
type IRenderComponent interface {
	Draw(*ebiten.Image, geometry.Point)
	GetDrawRectangle(geometry.Point) geometry.Rectangle
}

// RenderComponent implements the IRenderComponent interface
type RenderComponent struct {
	spriteSheetLibrary render.ISpriteSheetLibrary
	camera             render.ICamera
	name               string
}

// NewRenderComponent creates a IRenderComponent object
func NewRenderComponent(ssl render.ISpriteSheetLibrary, cam render.ICamera, name string) IRenderComponent {
	return &RenderComponent{
		spriteSheetLibrary: ssl,
		camera:             cam,
		name:               name,
	}
}

// Draw renders the component on the screen at the pointToDraw
func (r *RenderComponent) Draw(screen *ebiten.Image, pointToDraw geometry.Point) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(pointToDraw.X, pointToDraw.Y)
	ss := r.spriteSheetLibrary.GetSpriteSheet(r.name)
	ss.Draw(screen, r.camera, opts)
}

// GetDrawRectangle gives a rectangle which describes the RenderComponent at the position isoPoint
func (r *RenderComponent) GetDrawRectangle(isoPoint geometry.Point) geometry.Rectangle {
	ss := r.spriteSheetLibrary.GetSpriteSheet(r.name)
	return geometry.NewRectangle(
		float64(ss.Definition.FrameWidth),
		float64(ss.Definition.FrameHeight),
		isoPoint.X,
		isoPoint.Y,
	)
}
