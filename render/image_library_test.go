package render_test

import (
	"go_rts/render"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func Test_GivenImageLibrary_WhenGivenNameInDict_ThenReturnsImage(t *testing.T) {
	imageLibrary := render.NewImageLibraryFromPair("test", ebiten.NewImage(20, 20))

	img, err := imageLibrary.GetImage("test")

	if err != nil || img == nil {
		t.Errorf("Image returned not valid, error returned instead")
	}
}

func Test_GivenImageLibrary_WhenGivenNameNotInDict_ThenReturnsError(t *testing.T) {
	imageLibrary := render.NewImageLibraryFromPair("test", ebiten.NewImage(20, 20))

	img, err := imageLibrary.GetImage("invalid name")

	if err == nil || img != nil {
		t.Errorf("Should have returned error, key not in imageLibrary")
	}
}

func Test_GivenJSONDescribingImageLib_ThenReturnsImageLib(t *testing.T) {
	il := render.NewImageLibraryFromJSON("../assets/assets.json")

	img, err := il.GetImage("man")

	if img == nil || err != nil {
		t.Errorf("Should have returned image library with a man, and no error")
	}
}
