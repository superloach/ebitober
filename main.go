package main

import (
	"github.com/superloach/ebitober/2020/day01"
	"github.com/superloach/ebitober/2020/day02"
	"github.com/superloach/ebitober/ebitober"
)

func main() {
	ebitober.New(
		day01.New,
		day02.New,
	).Run()
}
