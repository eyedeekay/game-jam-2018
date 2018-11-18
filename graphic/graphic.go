package gamegraphic

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Graphic struct {
	colorslist  []Color
	sectionlist []section
	randomize   bool
	tags        []string
}

func (g Graphic) MinY() int {
	y := 32
	for _, s := range g.sectionlist {
		if s.y < y {
			y = s.y
		}
	}
	return y
}

func (g Graphic) Height() int {
	h := 0
	y := 0
	sh := 0
	for _, s := range g.sectionlist {
		if s.y > y || s.y == 0 {
			y = s.y
			h = s.y + s.h
			if h > sh {
				sh = h
			}
		}
	}
	r := sh - g.MinY()
	return r
}

func (g Graphic) MinX() int {
	x := 32
	for _, s := range g.sectionlist {
		if s.x < x {
			x = s.x
		}
	}
	return x
}

func (g Graphic) Width() int {
	w := 0
	x := 0
	sw := 0
	for _, s := range g.sectionlist {
		if s.x > x || s.x == 0 {
			x = s.x
			w = s.x + s.w
			if w >= sw {
				sw = w
			}
		}
	}
	r := sw - g.MinX()
	return r
}

func (g Graphic) ColorPick(tag ...string) Color {
	if len(tag) > 0 {
		ra := rand.Intn(len(tag))
		for _, key := range g.colorslist {
			if key.name == tag[ra] {
				return key
			}
		}
	}
	return g.ColorPick("black")
}

func (g Graphic) GetSectionList() []section {
	return g.sectionlist
}

func (g Graphic) GetColorList() []Color {
	return g.colorslist
}

func (g Graphic) GetRandomize() bool {
	return g.randomize
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (g Graphic) GetTags() []string {
	g.tags = unique(g.tags)
	return g.tags
}

func (g Graphic) HasTag(s ...string) bool {
	for _, t := range s {
		for _, u := range g.GetTags() {
			if t == u {
				return true
			}
		}
	}
	return false
}
func NewGraphics(colorfile, shapefile string, randomize bool, tags ...string) (*Graphic, error) {
	var g Graphic
	g.randomize = false
	colors, err := ioutil.ReadFile(colorfile)
	if err != nil {
		return nil, err
	}
	for _, key := range strings.Split(string(colors), ";") {
		if key != "" && key != "\n" {
			if c, err := newColor(key); err == nil {
				g.colorslist = append(g.colorslist, c)
			} else {
				return nil, err
			}
		}
	}
	s := strings.Split(strings.Split(strings.Split(shapefile, "/")[len(strings.Split(shapefile, "/"))-1], ".txt")[0], "_")
	for _, y := range s {
		g.tags = append(g.tags, y)
	}
	shapes, err := ioutil.ReadFile(shapefile)
	if err != nil {
		return nil, err
	}
	for _, key := range strings.Split(string(shapes), "\n") {
		if key != "" && key != "\n" {
			if s, err := newSection(key); err == nil {
				for _, val := range s.ColorSet() {
					g.tags = append(g.tags, val)
				}
				g.sectionlist = append(g.sectionlist, s)
			} else {
				return nil, err
			}
		}
	}

	log.Println("Loading graphic", colorfile, shapefile, "W:", g.Width(), "H:", g.Height())

	return &g, nil
}
