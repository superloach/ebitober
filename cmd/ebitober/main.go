package main

import (
	"github.com/superloach/ebitober"

	"github.com/superloach/ebitober/2020/day01"
	"github.com/superloach/ebitober/2020/day02"
)

func main() {
	ebitober.New(
		day01.New(13),
		day02.New(),
	).Run()
}
