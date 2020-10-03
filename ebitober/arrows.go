package ebitober

import "github.com/hajimehoshi/ebiten"

func (e *Ebitober) Arrows(s *ebiten.Image, aw, ah, w, h float64) {
	dio := &ebiten.DrawImageOptions{}
	dio.GeoM.Scale(-1, 1)
	dio.GeoM.Translate(aw, h-ah)
	_ = s.DrawImage(e.Arrow, dio)

	dio = &ebiten.DrawImageOptions{}
	dio.GeoM.Translate(w-aw, h-ah)
	_ = s.DrawImage(e.Arrow, dio)
}
