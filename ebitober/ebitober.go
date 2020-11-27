package ebitober

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Ebitober struct {
	Newers []Newer
	Days   []Day
	Day    int
	Banner *ebiten.Image
	Arrow  *ebiten.Image
}

func New(ns ...Newer) *Ebitober {
	if len(ns) == 0 {
		panic("need > 0 newers")
	}

	return &Ebitober{
		Newers: ns,
		Days:   make([]Day, 0, len(ns)),
		Day:    0,
		Banner: Image("banner"),
		Arrow:  Image("arrow"),
	}
}

func (e *Ebitober) Layout(ow, oh int) (int, int) {
	if ow > oh {
		return oh, oh
	}

	return ow, ow
}

func (e *Ebitober) Run() {
	rand.Seed(time.Now().Unix())

	ebiten.SetWindowResizable(true)

	go e.Load()

	err := ebiten.RunGame(e)
	if err != nil {
		panic(err)
	}
}
