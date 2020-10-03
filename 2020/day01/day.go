package day01

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"

	"github.com/superloach/ebitober/ebitober"
)

type Day struct {
	Fish   []*Fish
	Sprite *ebiten.Image
}

func New() ebitober.Day {
	const n = 13

	d := &Day{
		Fish:   make([]*Fish, 0, n),
		Sprite: ebitober.Image("day01/sprite"),
	}

	for i := 0; i < n; i++ {
		d.Fish = append(d.Fish, FishAt(-1, -1))
	}

	return d
}

func (d *Day) Info() string {
	return fmt.Sprintf("2020, Day 1 - Fish (%d of them)\n\nClick fish to remove them.\nClick elsewhere to make fish.", len(d.Fish))
}

func (d *Day) Tick(s *ebiten.Image, w, h float64) error {
	for _, fish := range d.Fish {
		switch {
		case fish.X < 0:
			fish.Vx = rand.Float64()*W + W/2
		case fish.X >= (w - W):
			fish.Vx = -(rand.Float64()*W + W/2)
		case fish.Y < 0:
			fish.Vy = rand.Float64()*H + H/2
		case fish.Y >= (h - H):
			fish.Vy = -(rand.Float64()*H + H/2)
		}

		fish.X += fish.Vx / 30
		fish.Y += fish.Vy / 30

		dio := &ebiten.DrawImageOptions{}

		dio.ColorM.RotateHue(fish.Col * math.Pi * 2)
		dio.ColorM.Scale(1, 1, 1, 0.9)

		if fish.Vx > 0 {
			dio.GeoM.Scale(-1, 1)
			dio.GeoM.Translate(W, 0)
		}
		dio.GeoM.Translate(fish.X, fish.Y)

		_ = s.DrawImage(d.Sprite, dio)
	}

	return nil
}

func (d *Day) Click(x, y, w, h float64, dur int) {
	if dur > 1 {
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
