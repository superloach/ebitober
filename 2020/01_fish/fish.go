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

func FishAt(x, y float64) *Fish {
	f := &Fish{
		X:   x,
		Y:   y,
		Vx:  rand.Float64()*W + W/2,
		Vy:  rand.Float64()*H + H/2,
		Col: rand.Float64(),
	}

	if rand.Float64() > 0.5 {
		f.Vx *= -1
	}

	if rand.Float64() > 0.5 {
		f.Vy *= -1
	}

	return f
}

func RandFish(bw, bh float64) *Fish {
	return FishAt(
		rand.Float64()*(bw-W),
		rand.Float64()*(bh-H),
	)
}

func (f *Fish) Update(screen *ebiten.Image) error {
	b := screen.Bounds()
	bw := float64(b.Dx())
	bh := float64(b.Dy())

	if f.X < -W || f.X >= bw || f.Y < -H || f.Y >= bh {
		f.Vx = rand.Float64()*W + W/2
		f.Vy = rand.Float64()*H + H/2
	}

	switch {
	case f.X < 0:
		f.Vx = rand.Float64()*W + W/2
	case f.X >= (bw - W):
		f.Vx = -(rand.Float64()*W + W/2)
	case f.Y < 0:
		f.Vy = rand.Float64()*H + H/2
	case f.Y >= (bh - H):
		f.Vy = -(rand.Float64()*H + H/2)
	}

	f.X += f.Vx / 30
	f.Y += f.Vy / 30

	return nil
}

func (f *Fish) Draw(screen *ebiten.Image) {
	dio := &ebiten.DrawImageOptions{}

	dio.ColorM.RotateHue(f.Col * math.Pi * 2)
	dio.ColorM.Scale(1, 1, 1, 0.9)

	if f.Vx > 0 {
		dio.GeoM.Scale(-1, 1)
		dio.GeoM.Translate(W, 0)
	}

	dio.GeoM.Translate(f.X, f.Y)

	_ = screen.DrawImage(Image, dio)
}
