package ebitober

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

//go:generate pkger

type Ebitober struct {
	Days []Day
	Day  int
}

func New(days ...Day) *Ebitober {
	if len(days) == 0 {
		panic("need > 0 days")
	}

	return &Ebitober{
		Days: days,
		Day:  0,
	}
}

func (e *Ebitober) Draw(screen *ebiten.Image) {
	day := e.Days[e.Day]

	day.Draw(screen)

	ebitenutil.DebugPrint(screen, "EBITOBER (by superloach)\n\n"+day.Info()+"\n\nPress <- or -> to change days.")
}

func (e *Ebitober) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		e.Day++
		e.Day %= len(e.Days)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		e.Day--
		e.Day += len(e.Days)
		e.Day %= len(e.Days)
	}

	err := e.Days[e.Day].Update(screen)
	if err != nil {
		err = fmt.Errorf("day %d update: %w", e.Day, err)
		return err
	}

	return nil
}

func (e *Ebitober) Layout(ow, oh int) (int, int) {
	return e.Days[e.Day].Layout(ow, oh)
}

func (e *Ebitober) Run() {
	ebiten.SetWindowResizable(true)

	err := ebiten.RunGame(e)
	if err != nil {
		panic(err)
	}
}
