package gamegraphic

import (
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	shape    string
	x        int
	y        int
	w        int
	h        int
	colorset []string
}

func (s section) X() int {
	return s.x
}

func (s section) Y() int {
	return s.y
}

func (s section) Width() int {
	return s.w
}

func (s section) Height() int {
	return s.h
}

func (s section) ColorSet() []string {
	var rst []string
	for _, key := range s.colorset {
		if key != "Scolor" && key != "scolor" && key != "" && key != " " && key != "\n" {
			k := strings.Replace(strings.Replace(key, "\n", "", -1), " ", "", -1)
			rst = append(rst, k)
		}
	}
	return rst
}

func newSection(s string) (section, error) {
	if s == "" {
		return section{
			shape:    "rect",
			x:        0,
			y:        0,
			w:        0,
			h:        0,
			colorset: []string{""},
		}, fmt.Errorf("Empty string passed to section initiator %s", s)
	}
	subs := strings.Split(s, ";")
	keys := strings.SplitN(subs[0], " ", 5)
	colors := strings.Split(strings.Replace(strings.Replace(subs[1], "scolor ", "", -1), "\n", "", -1), " ")
	x, err := strconv.Atoi(keys[1])
	if err != nil {
		return section{}, err
	}
	y, err := strconv.Atoi(keys[2])
	if err != nil {
		return section{}, err
	}
	w, err := strconv.Atoi(keys[3])
	if err != nil {
		return section{}, err
	}
	h, err := strconv.Atoi(keys[4])
	if err != nil {
		return section{}, err
	}
	return section{
		shape:    keys[0],
		x:        x,
		y:        y,
		w:        w,
		h:        h,
		colorset: colors,
	}, nil
}
