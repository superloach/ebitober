package ebitober

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (e *Ebitober) Update(s *ebiten.Image) error {
	sw, sh := s.Size()
	w, h := float64(sw), float64(sh)

	if e.Newers != nil {
		e.Loading(s, w, h)
		return nil
	}

	aw, ah := e.Arrow.Size()
	awf, ahf := float64(aw), float64(ah)

	day := e.Days[e.Day]

	err := day.Tick(s, w, h)
	if err != nil {
		err = fmt.Errorf("day %d update: %w", e.Day, err)
		return err
	}

	ebitenutil.DebugPrint(s, "\n\nEBITOBER (by superloach)\n\n"+day.Info())

	e.Arrows(s, w, h, awf, ahf)
	e.Inputs(day, w, h)

	return nil
}
