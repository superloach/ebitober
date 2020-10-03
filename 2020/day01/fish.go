package day01

import (
	"math/rand"
)

const (
	W = 100
	H = 64
)

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

func RandFish(w, h float64) *Fish {
	return FishAt(
		rand.Float64()*(w-W),
		rand.Float64()*(h-H),
	)
}
