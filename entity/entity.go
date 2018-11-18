package gameentity

import (
	//	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	graphic "../graphic"
)

type Entity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewEntityFromGraphic(g graphic.Graphic, x, y float32) (*Entity, error) {
	texture := GenerateTexture(x, y, g)
	rendercomponent := common.RenderComponent{
		Drawable: texture,
	}
	spacecomponent := common.SpaceComponent{
		Position: engo.Point{x, y},
	}
	return &Entity{
		BasicEntity:     ecs.NewBasic(),
		SpaceComponent:  spacecomponent,
		RenderComponent: rendercomponent,
	}, nil
}
