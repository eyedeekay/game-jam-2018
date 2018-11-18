package gamegraphic

import (
	"log"
	"math/rand"
	"path/filepath"
)

type GraphicsFactory struct {
	graphics []Graphic
}

func (g GraphicsFactory) FromTags(s ...string) Graphic {
	var x []Graphic
	if len(s) > 0 {
		for _, key := range g.graphics {
			if key.HasTag(s...) {
				x = append(x, key)
			}
		}
	}
	gen := 1
	if len(x) > 1 {
		gen = len(x) - 1
		val := rand.Intn(gen)
		return x[val]
	}
	return x[0]

}

func NewGraphicsFactory(assetpath, palette, skelfolder, configfile string) (*GraphicsFactory, error) {
	log.Println("Loading graphics factory")
	var g GraphicsFactory
	dirs, err := filepath.Glob(filepath.Join("./", assetpath, skelfolder, "/*.txt"))
	if err != nil {
		return nil, err
	}
	for i, key := range dirs {
		log.Println("loading", i, key)
		gt, err := NewGraphics(
			filepath.Join(assetpath, palette),
			filepath.Join(key),
			false,
		)
		if err != nil {
			return nil, err
		}
		g.graphics = append(g.graphics, *gt)
	}
	log.Println("Loaded graphics factory", len(dirs), "structures loaded")
	return &g, nil
}
