package fish

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Day struct {
	Fish []*Fish
}

func New(n int) *Day {
	d := &Day{
		Fish: make([]*Fish, n),
	}

	for i := 0; i < n; i++ {
		d.Fish[i] = RandFish()
	}

	return d
}

func (d *Day) Info() string {
	return "10/01/2020 - Fish"
}

func (d *Day) Draw(screen *ebiten.Image) {
	for _, fish := range d.Fish {
		fish.Draw(screen)
	}
}

func (d *Day) Update(screen *ebiten.Image) error {
	for i, fish := range d.Fish {
		err := fish.Update(screen)
		if err != nil {
			return fmt.Errorf("fish %d update: %w", i, err)
		}
	}

	return nil
}

func (d *Day) Layout(ow, oh int) (int, int) {
	return ow, oh
}
