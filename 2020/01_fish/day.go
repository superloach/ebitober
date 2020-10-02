package fish

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Day struct {
	N    int
	Init bool
	Fish []*Fish
}

func (d *Day) Info() string {
	return fmt.Sprintf("10/01/2020 - (%d) Fish", len(d.Fish))
}

func (d *Day) Update(screen *ebiten.Image) error {
	b := screen.Bounds()
	bw := float64(b.Dx())
	bh := float64(b.Dy())

	if !d.Init {
		d.Fish = append(d.Fish, RandFish(bw, bh))

		if len(d.Fish) == d.N {
			d.Init = true
		}
	}

	for i, fish := range d.Fish {
		err := fish.Update(screen)
		if err != nil {
			return fmt.Errorf("fish %d update: %w", i, err)
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		fcx, fcy := float64(cx), float64(cy)

		sel := -1

		for i, fish := range d.Fish {
			if fcx >= fish.X && fcx <= fish.X+W && fcy >= fish.Y && fcy <= fish.Y+H {
				sel = i
			}
		}

		if sel != -1 {
			l := len(d.Fish) - 1
			d.Fish[sel] = d.Fish[l]
			d.Fish[l] = nil
			d.Fish = d.Fish[:l]
		} else {
			d.Fish = append(d.Fish, FishAt(fcx, fcy))
		}
	}

	return nil
}

func (d *Day) Draw(screen *ebiten.Image) {
	for _, fish := range d.Fish {
		fish.Draw(screen)
	}
}

func (d *Day) Layout(ow, oh int) (int, int) {
	return ow, oh
}
