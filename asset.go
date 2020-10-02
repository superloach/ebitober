package ebitober

import (
	"image/png"
	"net/http"

	"github.com/hajimehoshi/ebiten"
)

//go:generate pkger

func Asset(f http.File, err error) *ebiten.Image {
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	img, err := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	return img
}
