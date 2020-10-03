package day02

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten"

	"github.com/superloach/ebitober/ebitober"
)

type Day struct {
	Wisp *ebiten.Image
	Wx   float64

	Rock   *ebiten.Image
	Rx, Ry float64

	Score, Hi uint
}

func New() ebitober.Day {
	return &Day{
		Wisp: ebitober.Image("day02/wisp"),
		Wx:   0,

		Rock: ebitober.Image("day02/rock"),
		Rx:   123,
		Ry:   0,

		Score: 0,
		Hi:    10,
	}
}

func (d *Day) Info() string {
	return fmt.Sprintf("2020, Day 2 - Wisp (score %d, hi %d)\n\nA single rock falls at a time.\nClick to move your wisp and avoid it.", d.Score, d.Hi)
}

func (d *Day) Tick(s *ebiten.Image, w, h float64) error {
	mh := 2 * h / 3

	ww, wh := d.Wisp.Size()
	fww, fwh := float64(ww), float64(wh)

	rw, rh := d.Rock.Size()
	frw, frh := float64(rw), float64(rh)

	d.Ry += float64(d.Score) + 5
	d.Rx += (d.Wx - d.Rx) / (200 / float64(d.Score+5))

	switch {
	case d.Ry > h:
		d.Score++
		d.Ry = 0
		d.Rx = rand.Float64() * w
	case d.Rx > w:
		d.Rx = w
	}

	if d.Wx > d.Rx-frw/2 && d.Wx < d.Rx+frw/2 && mh > d.Ry-fwh/2 && mh < d.Ry+fwh/2 {
		d.Score = 0
		d.Ry = 0
		d.Rx = rand.Float64() * w
	}

	dio := &ebiten.DrawImageOptions{}
	dio.GeoM.Translate(d.Wx-fww/2, mh-fwh/2)
	_ = s.DrawImage(d.Wisp, dio)
	_ = s.DrawImage(d.Wisp, dio)

	dio = &ebiten.DrawImageOptions{}
	dio.GeoM.Translate(d.Rx-frw/2, d.Ry-frh/2)
	_ = s.DrawImage(d.Rock, dio)

	if d.Score > d.Hi {
		d.Hi = d.Score
	}

	return nil
}

func (d *Day) Click(x, y, w, h float64, dur int) {
	d.Wx = x
}
