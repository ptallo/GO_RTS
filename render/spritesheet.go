package render

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewSpriteSheetMap() map[string]*SpriteSheet {
	man, _ := NewSpriteSheet("units", "man")
	woman, _ := NewSpriteSheet("units", "woman")
	grass, _ := NewSpriteSheet("tiles", "grass")
	water, _ := NewSpriteSheet("tiles", "water")
	return map[string]*SpriteSheet{
		"man":   &man,
		"woman": &woman,
		"grass": &grass,
		"water": &water,
	}
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

func NewSpriteSheet(assetDir, name string) (SpriteSheet, error) {
	var ss SpriteSheet

	jsonPath := fmt.Sprintf("./assets/%v/%v.json", assetDir, name)
	ssd, err := NewSpriteSheetDefinitionFromJSON(jsonPath)

	if err != nil {
		return ss, err
	}

	imgPath := fmt.Sprintf("./assets/%v/%v.png", assetDir, name)
	img, err := NewImageFromPath(imgPath)

	if err != nil {
		return ss, err
	}

	ss = SpriteSheet{
		Image:      img,
		Definition: ssd,
	}

	return ss, nil
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
