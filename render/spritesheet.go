package render

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheetLibrary struct {
	Library map[string]*SpriteSheet
}

func (ssl *SpriteSheetLibrary) GetSpriteSheet(name string) *SpriteSheet {
	return ssl.Library[name]
}

type SpriteSheet struct {
	Image      *ebiten.Image
	Definition SpriteSheetDefinition
}

type SpriteSheetDefinition struct {
	Width       int
	Height      int
	FrameWidth  int
	FrameHeight int
}

func NewSpriteSheet(assetDir, name string) (*SpriteSheet, error) {
	jsonPath := fmt.Sprintf("./assets/%v/%v.json", assetDir, name)
	ssd, err := NewSpriteSheetDefinitionFromJSON(jsonPath)

	if err != nil {
		return nil, err
	}

	imgPath := fmt.Sprintf("./assets/%v/%v.png", assetDir, name)
	img, err := NewImageFromPath(imgPath)

	if err != nil {
		return nil, err
	}

	return &SpriteSheet{
		Image:      img,
		Definition: ssd,
	}, nil
}

func NewSpriteSheetDefinitionFromJSON(path string) (SpriteSheetDefinition, error) {
	var ssd SpriteSheetDefinition

	// Read file into data array
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ssd, err
	}

	// Transform data array into JSON structure
	err = json.Unmarshal(data, &ssd)
	if err != nil {
		return ssd, err
	}

	return ssd, err
}

func (ss *SpriteSheet) Draw(screen *ebiten.Image, camera *Camera, opts *ebiten.DrawImageOptions) {
	rect := image.Rect(0, 0, ss.Definition.FrameWidth, ss.Definition.FrameHeight)
	img := ss.Image.SubImage(rect).(*ebiten.Image)
	camera.DrawImage(screen, img, opts)
}
