package ebitober

import (
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/markbates/pkger"
)

//go:generate pkger

func Img(name string) *ebiten.Image {
	pkger.Include("/assets")

	f, err := pkger.Open("/assets/" + name + ".png")
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
