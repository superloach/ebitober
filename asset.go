package ebitober

import (
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/markbates/pkger"
)

//go:generate pkger

func Asset(path string) *ebiten.Image {
	f, err := pkger.Open(path)
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
