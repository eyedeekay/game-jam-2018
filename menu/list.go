package menu

import (
	"image/color"
	//"github.com/EngoEngine/ecs"
	//"github.com/EngoEngine/engo"
	//"github.com/EngoEngine/engo/common"
)

var (
	brown = color.RGBA{70, 0, 0, 255}
	light = color.RGBA{153, 204, 153, 255}
	hide  = color.RGBA{0, 0, 0, 0}
)

type Listing struct {
	Elements []*Div
	index    int
	hidden   bool
}

func (l *Listing) Hide() {
	for index := range l.Elements {
		l.Elements[index].SwitchColor(hide, hide)
	}
	l.hidden = true
}

func (l *Listing) Show() {
	for index := range l.Elements {
		l.Elements[index].SwitchColor(light, brown)
	}
	l.Elements[0].SwitchColor(brown, light)
	l.hidden = false
}

func (l *Listing) SwitchIndexUp() {
	if l.index == len(l.Elements)-1 {
		l.index = 0
	} else {
		l.index++
	}
	for index := range l.Elements {
		if l.index == index {
			l.Elements[index].SwitchColor(brown, light)
		} else {
			l.Elements[index].SwitchColor(light, brown)
		}
	}
}

func (l *Listing) SwitchIndexDown() {
	if l.index == 0 {
		l.index = len(l.Elements) - 1
	} else {
		l.index--
	}
	for index := range l.Elements {
		if l.index == index {
			l.Elements[index].SwitchColor(brown, light)
		} else {
			l.Elements[index].SwitchColor(light, brown)
		}
	}
}

func NewListing(topx, topy float32, msg []string) *Listing {
	listing := &Listing{}
	for index, txt := range msg {
		listing.Elements = append(listing.Elements, NewColorDiv(txt, 1, 24, topx, topy+(float32(index)*44), light, brown))
	}
	listing.index = len(listing.Elements) - 1
	listing.SwitchIndexUp()
	return listing
}
