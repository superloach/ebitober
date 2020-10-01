package fish

import (
	"image/png"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/markbates/pkger"
)

const (
	W = 100
	H = 64
)

var Cross = math.Sqrt(W*W + H*H)

var Image = func() *ebiten.Image {
	f, err := pkger.Open("/assets/01_fish.png")
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
}()
