package main

import (
	"github.com/superloach/ebitober"

	"github.com/superloach/ebitober/2020/day01"
)

func main() {
	ebitober.New(
		&day01.Day{N: 13},
	).Run()
}
