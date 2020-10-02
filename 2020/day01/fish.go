package day01

import (
	"math/rand"
	"time"
)

var _ = func() bool {
	rand.Seed(time.Now().Unix())
	return true
}()

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
