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

func CloneImage(origImage *ebiten.Image) *ebiten.Image {
	w, h := origImage.Size()
	img := ebiten.NewImage(w, h)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	img.DrawImage(origImage, op)
	return img
}
