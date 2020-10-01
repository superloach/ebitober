package ebitober

import "github.com/hajimehoshi/ebiten"

type Day interface {
	ebiten.Game

	Draw(*ebiten.Image)
	Info() string
}
