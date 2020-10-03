package ebitober

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Ebitober struct {
	Days  []Day
	Day   int
	Arrow *ebiten.Image
}

func New(days ...Day) *Ebitober {
	if len(days) == 0 {
		panic("need > 0 days")
	}

	return &Ebitober{
		Days:  days,
		Day:   0,
		Arrow: Img("arrow"),
	}
}

func (e *Ebitober) Layout(ow, oh int) (int, int) {
	return ow, oh
}

func (e *Ebitober) Run() {
	rand.Seed(time.Now().Unix())

	ebiten.SetWindowResizable(true)

	err := ebiten.RunGame(e)
	if err != nil {
		panic(err)
	}
}
