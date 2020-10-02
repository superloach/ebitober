package main

import (
	"github.com/superloach/ebitober"

	fish "github.com/superloach/ebitober/2020/01_fish"
)

func main() {
	ebitober.New(
		&fish.Day{N: 13},
	).Run()
}
