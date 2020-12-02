package render

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteSheetLibrary is a wrapper around a map[string]*SpriteSheet
type SpriteSheetLibrary struct {
	Library map[string]*SpriteSheet
}

// GetSpriteSheet gets a spritesheet from the SpriteSheetLibrary
func (ssl *SpriteSheetLibrary) GetSpriteSheet(name string) *SpriteSheet {
	return ssl.Library[name]
}

// SpriteSheet is a way to define a spritesheet given a definition and an Image
type SpriteSheet struct {
	Image      *ebiten.Image
	Definition SpriteSheetDefinition
}

// SpriteSheetDefinition description of a SpriteSheet image which tells us how to use the spritesheet
type SpriteSheetDefinition struct {
	Width       int
	Height      int
	FrameWidth  int
	FrameHeight int
}

// NewSpriteSheet is a shortcut to load a spritesheet given the name of the asset dir and the name of the asset
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

// NewSpriteSheetDefinitionFromJSON deserializes a spritesheet definition from a json file given the path
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

// Draw will draw a SpriteSheet on a given screen and camera
func (ss *SpriteSheet) Draw(screen *ebiten.Image, camera *Camera, opts *ebiten.DrawImageOptions) {
	rect := image.Rect(0, 0, ss.Definition.FrameWidth, ss.Definition.FrameHeight)
	img := ss.Image.SubImage(rect).(*ebiten.Image)
	camera.DrawImage(screen, img, opts)
}
