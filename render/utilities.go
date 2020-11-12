package render

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImageFromPath(path string) (*ebiten.Image, error) {
	reader, err := os.Open(path)
	defer reader.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
