package day01

import (
	"math/rand"

	"github.com/markbates/pkger"
	"github.com/superloach/ebitober"
)

const (
	W = 100
	H = 64
)

var Image = ebitober.Asset(pkger.Open("/2020/day01/assets/fish.png"))

type Fish struct {
	X, Y   float64
	Vx, Vy float64
	Col    float64
}

func FishAt(x, y float64) *Fish {
	f := &Fish{
		X:   x,
		Y:   y,
		Vx:  rand.Float64()*W + W/2,
		Vy:  rand.Float64()*H + H/2,
		Col: rand.Float64(),
	}

	if rand.Float64() > 0.5 {
		f.Vx *= -1
	}

	if rand.Float64() > 0.5 {
		f.Vy *= -1
	}

	return f
}

func RandFish(bw, bh float64) *Fish {
	return FishAt(
		rand.Float64()*(bw-W),
		rand.Float64()*(bh-H),
	)
}
