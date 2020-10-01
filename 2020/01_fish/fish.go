package fish

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var _ = func() bool {
	rand.Seed(time.Now().Unix())
	return true
}()

type Fish struct {
	X, Y   float64
	Vx, Vy float64
	Col    float64
}

func RandFish() *Fish {
	return &Fish{
		X:   rand.Float64() * 0.8,
		Y:   rand.Float64() * 0.9,
		Vx:  rand.Float64(),
		Vy:  rand.Float64(),
		Col: rand.Float64(),
	}
}

func (f *Fish) Update(screen *ebiten.Image) error {
	if f.X < 0 || f.X > 0.8 {
		f.Vx *= -1
	}

	if f.Y < 0 || f.Y > 0.9 {
		f.Vy *= -1
	}

	f.X += f.Vx / 120
	f.Y += f.Vy / 120

	return nil
}

func (f *Fish) Scale(bw, bh float64) float64 {
	return math.Sqrt(bw*bw+bh*bh) / (5 * Cross)
}

func (f *Fish) Draw(screen *ebiten.Image) {
	b := screen.Bounds()
	bw := float64(b.Dx())
	bh := float64(b.Dy())

	scale := f.Scale(bw, bh)

	dio := &ebiten.DrawImageOptions{}

	dio.ColorM.RotateHue(f.Col * math.Pi * 2)
	dio.ColorM.Scale(1, 1, 1, 0.9)

	dio.GeoM.Scale(scale, scale)

	if f.Vx > 0 {
		dio.GeoM.Scale(-1, 1)
		dio.GeoM.Translate(bw/5, 0)
	}

	dio.GeoM.Translate(f.X*bw, f.Y*bh)

	_ = screen.DrawImage(Image, dio)
}
