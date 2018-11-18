package gamegraphic

import (
	"fmt"
	icolor "image/color"
	"strconv"
	"strings"
)

type Color struct {
	name string
	r    int
	g    int
	b    int
	a    int
}

func (c Color) NRGBA() *icolor.NRGBA {
	return &icolor.NRGBA{
		R: uint8(c.r),
		G: uint8(c.g),
		B: uint8(c.b),
		A: uint8(c.a),
	}
}

func newColor(s string) (Color, error) {
	if test := strings.Replace(s, "\n", "", -1); test == "" {
		return Color{
			name: "ERRRED",
			r:    0,
			g:    0,
			b:    0,
			a:    0,
		}, nil
	}
	v := strings.Split(s, " ")
	if len(v) == 6 {
		red, err := strconv.Atoi(v[2])
		if err != nil {
			return Color{}, err
		}
		green, err := strconv.Atoi(v[3])
		if err != nil {
			return Color{}, err
		}
		blue, err := strconv.Atoi(v[4])
		if err != nil {
			return Color{}, err
		}
		alpha, err := strconv.Atoi(v[5])
		if err != nil {
			return Color{}, err
		}
		return Color{
			name: v[1],
			r:    red,
			g:    green,
			b:    blue,
			a:    alpha,
		}, nil
	} else {
		return Color{}, fmt.Errorf("Error loading color %s len %d ", s, len(v))
	}
}
