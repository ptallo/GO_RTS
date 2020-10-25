package render

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImageFromPath(path string) *ebiten.Image {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
