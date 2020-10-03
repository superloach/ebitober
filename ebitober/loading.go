package ebitober

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (e *Ebitober) Load() {
	dur := 1 * time.Second / time.Duration(len(e.Newers))

	for _, n := range e.Newers {
		time.Sleep(dur)
		e.Days = append(e.Days, n())
	}

	time.Sleep(dur)
	e.Newers = nil
}

func (e *Ebitober) Loading(s *ebiten.Image, w, h float64) {
	bw, bh := e.Banner.Size()
	bwf, bhf := float64(bw), float64(bh)

	scale := fit(w, h, bwf, bhf)

	mx := w/2 - scale*bwf/2
	my := h/2 - scale*bhf/2

	dio := &ebiten.DrawImageOptions{}
	dio.GeoM.Scale(scale, scale)
	dio.GeoM.Translate(mx, my)

	r := e.Banner.Bounds()
	r.Max.X = (r.Max.X * (len(e.Newers) + 1)) / (len(e.Days) + 1)

	_ = s.DrawImage(e.Banner, dio)

	ebitenutil.DebugPrint(s, fmt.Sprintf("\n\nEBITOBER (by superloach)\n\nloading %d/%d days", len(e.Days), len(e.Newers)))
}

func fit(w, h, bw, bh float64) float64 {
	if w/h > bw/bh {
		return h / bh
	}

	return w / bw
}
