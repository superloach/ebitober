package ebitober

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	aW = 100
	aH = 100
)

func (e *Ebitober) Update(s *ebiten.Image) error {
	sw, sh := s.Size()
	w, h := float64(sw), float64(sh)

	day := e.Days[e.Day]

	err := day.Tick(s, w, h)
	if err != nil {
		err = fmt.Errorf("day %d update: %w", e.Day, err)
		return err
	}

	ebitenutil.DebugPrint(s, "\n\nEBITOBER (by superloach)\n\n"+day.Info())

	dio := &ebiten.DrawImageOptions{}
	dio.GeoM.Scale(-1, 1)
	dio.GeoM.Translate(aW, h-aH)
	_ = s.DrawImage(e.Arrow, dio)

	dio = &ebiten.DrawImageOptions{}
	dio.GeoM.Translate(w-aW, h-aH)
	_ = s.DrawImage(e.Arrow, dio)

	for _, btn := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonMiddle,
		ebiten.MouseButtonRight,
	} {
		if inpututil.IsMouseButtonJustPressed(btn) {
			cx, cy := ebiten.CursorPosition()
			x, y := float64(cx), float64(cy)

			e.Click(day, x, y, w, h)
		}
	}

	for _, tid := range inpututil.JustPressedTouchIDs() {
		tx, ty := ebiten.TouchPosition(tid)
		x, y := float64(tx), float64(ty)

		e.Click(day, x, y, w, h)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		e.Click(nil, -1, 0, w, h)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		e.Click(nil, 0, -1, w, h)
	}

	return nil
}

func (e *Ebitober) Click(day Day, x, y, w, h float64) {
	switch {
	case x == -1, x < aW && h-y < aH:
		fmt.Println("prev")
		e.Day++
		e.Day %= len(e.Days)
	case y == -1, w-x < aW && h-y < aH:
		fmt.Println("next")
		e.Day--
		e.Day += len(e.Days)
		e.Day %= len(e.Days)
	default:
		day.Click(x, y, w, h)
	}
}
