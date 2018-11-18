package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	//"log"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
	common.RenderComponent
}
