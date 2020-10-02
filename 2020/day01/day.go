package day01

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Day struct {
	N    int
	Init bool
	Fish []*Fish
}

func (d *Day) Info() string {
	return fmt.Sprintf("2020, Day 1 - Fish (%d of them)", len(d.Fish))
}

func (d *Day) Update(screen *ebiten.Image) error {
	b := screen.Bounds()
	bw := float64(b.Dx())
	bh := float64(b.Dy())

	if !d.Init {
		d.Fish = append(d.Fish, RandFish(bw, bh))

		if len(d.Fish) == d.N {
			d.Init = true
		}
	}

	for _, fish := range d.Fish {
		if fish.X < -W || fish.X >= bw || fish.Y < -H || fish.Y >= bh {
			fish.Vx = rand.Float64()*W + W/2
			fish.Vy = rand.Float64()*H + H/2
		}

		switch {
		case fish.X < 0:
			fish.Vx = rand.Float64()*W + W/2
		case fish.X >= (bw - W):
			fish.Vx = -(rand.Float64()*W + W/2)
		case fish.Y < 0:
			fish.Vy = rand.Float64()*H + H/2
		case fish.Y >= (bh - H):
			fish.Vy = -(rand.Float64()*H + H/2)
		}

		fish.X += fish.Vx / 30
		fish.Y += fish.Vy / 30
	}

	for _, btn := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonMiddle,
		ebiten.MouseButtonRight,
	} {
		if inpututil.IsMouseButtonJustPressed(btn) {
			cx, cy := ebiten.CursorPosition()
			d.Click(bw, bh, float64(cx), float64(cy))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		d.Click(bw, bh, -1, -1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		d.Click(bw, bh, -2, -2)
	}

	for _, tid := range inpututil.JustPressedTouchIDs() {
		tx, ty := ebiten.TouchPosition(tid)
		d.Click(bw, bh, float64(tx), float64(ty))
	}

	return nil
}

func (d *Day) Click(bw, bh, x, y float64) {
	if x == -1 && y == -1 {
		d.Fish = append(d.Fish, RandFish(bw, bh))

		return
	}

	if x == -2 && y == -2 {
		l := len(d.Fish) - 1

		d.Fish[l] = nil
		d.Fish = d.Fish[:l]

		return
	}

	sel := -1

	for i, fish := range d.Fish {
		if x >= fish.X && x <= fish.X+W && y >= fish.Y && y <= fish.Y+H {
			sel = i
		}
	}

	if sel != -1 {
		l := len(d.Fish) - 1

		d.Fish[sel] = d.Fish[l]
		d.Fish[l] = nil
		d.Fish = d.Fish[:l]

		return
	}

	d.Fish = append(d.Fish, FishAt(x, y))
}

func (d *Day) Draw(screen *ebiten.Image) {
	for _, fish := range d.Fish {
		dio := &ebiten.DrawImageOptions{}

		dio.ColorM.RotateHue(fish.Col * math.Pi * 2)
		dio.ColorM.Scale(1, 1, 1, 0.9)

		if fish.Vx > 0 {
			dio.GeoM.Scale(-1, 1)
			dio.GeoM.Translate(W, 0)
		}

		dio.GeoM.Translate(fish.X, fish.Y)

		_ = screen.DrawImage(Image, dio)
	}
}

func (d *Day) Layout(ow, oh int) (int, int) {
	return ow, oh
}
