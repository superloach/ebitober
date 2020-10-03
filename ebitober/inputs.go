package ebitober

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (e *Ebitober) Inputs(day Day, w, h float64) {
	for _, btn := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonMiddle,
		ebiten.MouseButtonRight,
	} {
		if ebiten.IsMouseButtonPressed(btn) {
			cx, cy := ebiten.CursorPosition()
			x, y := float64(cx), float64(cy)

			dur := inpututil.MouseButtonPressDuration(btn)

			e.Click(day, x, y, w, h, dur)
		}
	}

	for _, tid := range ebiten.TouchIDs() {
		tx, ty := ebiten.TouchPosition(tid)
		x, y := float64(tx), float64(ty)

		dur := inpututil.TouchPressDuration(tid)

		e.Click(day, x, y, w, h, dur)
	}
}

func (e *Ebitober) Click(day Day, x, y, w, h float64, dur int) {
	aw, ah := e.Arrow.Size()
	awf, ahf := float64(aw), float64(ah)

	switch {
	case dur == 1 && x < awf && h-y < ahf:
		fmt.Println("prev")
		e.Day++
		e.Day %= len(e.Days)
	case dur == 1 && w-x < awf && h-y < ahf:
		fmt.Println("next")
		e.Day--
		e.Day += len(e.Days)
		e.Day %= len(e.Days)
	default:
		day.Click(x, y, w, h, dur)
	}
}
