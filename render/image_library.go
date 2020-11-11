package render

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageLibrary struct {
	stringToImageMap map[string]*ebiten.Image
}

type SpriteSheetDefinition struct {
	Name        string
	Width       int
	Height      int
	FrameWidth  int
	FrameHeight int
}

type DefinitionLibrary struct {
	Units []SpriteSheetDefinition
	Tiles []SpriteSheetDefinition
}

func NewImageLibrary() ImageLibrary {
	m := map[string]*ebiten.Image{
		"man":   NewImageFromPath("./assets/units/Man.png"),
		"woman": NewImageFromPath("./assets/units/Woman.png"),
		"grass": NewImageFromPath("./assets/tiles/grass.png"),
		"water": NewImageFromPath("./assets/tiles/water.png"),
	}

	return ImageLibrary{
		stringToImageMap: m,
	}
}

func NewImageLibraryFromPair(name string, img *ebiten.Image) ImageLibrary {
	m := map[string]*ebiten.Image{
		name: img,
	}

	return ImageLibrary{
		stringToImageMap: m,
	}
}

func NewImageLibraryFromJSON(path string) ImageLibrary {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var dl DefinitionLibrary
	err = json.Unmarshal(dat, &dl)

	if err != nil {
		panic(err)
	}

	// TODO: Create image lib from the definition library

	return NewImageLibrary()
}

func (il *ImageLibrary) GetImage(name string) (*ebiten.Image, error) {
	img, ok := il.stringToImageMap[name]

	if !ok {
		return nil, fmt.Errorf("Image name %v doesn't exist in stringToImageMap", name)
	}

	return img, nil
}
