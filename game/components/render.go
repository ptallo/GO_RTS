package components

import (
	"go_rts/geometry"
	"go_rts/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type IRenderComponent interface {
	Draw(*ebiten.Image, geometry.Point)
	GetDrawRectangle(geometry.Point) geometry.Rectangle
}

type RenderComponent struct {
	spriteSheetLibrary render.ISpriteSheetLibrary
	camera             render.ICamera
	name               string
}

func NewRenderComponent(ssl render.ISpriteSheetLibrary, cam render.ICamera, name string) IRenderComponent {
	return &RenderComponent{
		spriteSheetLibrary: ssl,
		camera:             cam,
		name:               name,
	}
}

func (r *RenderComponent) Draw(screen *ebiten.Image, pointToDraw geometry.Point) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(pointToDraw.X, pointToDraw.Y)
	ss := r.spriteSheetLibrary.GetSpriteSheet(r.name)
	ss.Draw(screen, r.camera, opts)
}

func (r *RenderComponent) GetDrawRectangle(isoPoint geometry.Point) geometry.Rectangle {
	ss := r.spriteSheetLibrary.GetSpriteSheet(r.name)
	return geometry.NewRectangle(
		float64(ss.Definition.FrameWidth),
		float64(ss.Definition.FrameHeight),
		isoPoint.X,
		isoPoint.Y,
	)
}
