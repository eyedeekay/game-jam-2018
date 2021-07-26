package main

import (
	"flag"
	"log"

	"github.com/kbinani/screenshot"
)

var (
	height   = flag.Int("h", getheight(), "Height of the window in Pixels")
	width    = flag.Int("w", getwidth(), "Width of the window in Pixels")
	maximize = flag.Bool("m", true, "Start the window Maximized")
)

func getheight() int {
	bounds := screenshot.GetDisplayBounds(0)

	y := 768
	if bounds.Dy() <= 768 {
		y = 600
	}
	if *maximize {
		y = bounds.Dy() - 64
	}
	log.Println("Height:", y)
	return y
}

func getwidth() int {
	bounds := screenshot.GetDisplayBounds(0)

	x := 1024
	if bounds.Dx() <= 1024 {
		x = 800
	}
	if bounds.Dx() <= 1280 {
		x = 1024
	}
	if *maximize {
		x = bounds.Dx()
	}
	log.Println("Width:", x)
	return x
}
