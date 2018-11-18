package gameentity

import (
	"image"

	"engo.io/engo/common"

	graphic "../graphic"
)

func GenerateTexture(x, y float32, g graphic.Graphic) common.Texture {
	nrgba := image.NewNRGBA(image.Rect(
		g.MinX(),
		g.MinY(),
		g.MinX()+g.Width(),
		g.MinY()+g.Height(),
	))
	var c graphic.Color
	for _, key := range g.GetSectionList() {
		if !g.GetRandomize() {
			c = g.ColorPick(key.ColorSet()...)
		}
		for xx := key.X(); xx <= key.X()+key.Width(); xx++ {
			for yy := key.Y(); yy <= key.Y()+key.Height(); yy++ {
				if g.GetRandomize() {
					nrgba.Set(xx, yy, g.ColorPick(key.ColorSet()...).NRGBA())
				} else {
					nrgba.Set(xx, yy, c.NRGBA())
				}
			}
		}
	}
	object := common.NewImageObject(nrgba)
	texture := common.NewTextureSingle(object)
	return texture
}
