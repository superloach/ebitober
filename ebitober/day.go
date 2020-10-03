package ebitober

import (
	"github.com/hajimehoshi/ebiten"
)

type Day interface {
	Tick(s *ebiten.Image, w, h float64) error
	Click(x, y, w, h float64, dur int)
	Info() string
}

type Newer func() Day
