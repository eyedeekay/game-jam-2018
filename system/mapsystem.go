package systems

import (
	"fmt"
	"io/ioutil"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"../entity"
	graphic "../graphic"
	"github.com/mattn/anko/vm"
)

type MapSystem struct {
	vmEnv          *vm.Env
	World          *ecs.World
	MouseTracker   MouseTracker
	GraphicFactory *graphic.GraphicsFactory
	scriptpath     string
	script         string
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*MapSystem) Remove(ecs.BasicEntity) {

}

func (m *MapSystem) LoadEntity(x, y int, tags string) {
	brick, err := gameentity.NewEntityFromGraphic(
		m.GraphicFactory.FromTags(tags),
		float32(x),
		float32(y),
	)
	if err != nil {
		panic("Unable to load texture: " + err.Error())
	}
	for _, system := range m.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&brick.BasicEntity, &brick.RenderComponent, &brick.SpaceComponent)
		case *ScaleSystem:
			sys.Add(&brick.BasicEntity, &brick.RenderComponent, &brick.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&brick.BasicEntity,
				&m.MouseTracker.MouseComponent,
				&brick.SpaceComponent,
				&brick.RenderComponent,
			)
		}
	}
}

func (m *MapSystem) LoadWall(x, y float32) {
	brick, err := gameentity.NewEntityFromGraphic(
		m.GraphicFactory.FromTags("wall"),
		x,
		y,
	)
	if err != nil {
		panic("Unable to load texture: " + err.Error())
	}
	for _, system := range m.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&brick.BasicEntity, &brick.RenderComponent, &brick.SpaceComponent)
		case *ScaleSystem:
			sys.Add(&brick.BasicEntity, &brick.RenderComponent, &brick.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&brick.BasicEntity,
				&m.MouseTracker.MouseComponent,
				&brick.SpaceComponent,
				&brick.RenderComponent,
			)
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (m *MapSystem) Update(dt float32) {
	if engo.Input.Button("AddWall").JustPressed() {
		fmt.Println("The gamer pressed F1")
		m.LoadWall(m.MouseTracker.MouseX, m.MouseTracker.MouseY)
	}
	if engo.Input.Button("Exit").JustPressed() {
		engo.Exit()
	}
}

// New is the initialisation of the System
func (m *MapSystem) New(w *ecs.World, g *graphic.GraphicsFactory, scriptpath string) {
	log.Println("MapSystem was added to the Scene")
	m.World = w
	m.GraphicFactory = g
	m.MouseTracker.BasicEntity = ecs.NewBasic()
	m.MouseTracker.MouseComponent = common.MouseComponent{Track: true}
	m.scriptpath = scriptpath
	tmpscript, err := ioutil.ReadFile(m.scriptpath)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
	m.script = string(tmpscript)
	m.vmEnv = vm.NewEnv()
	err = m.vmEnv.Define("println", fmt.Println)
	err = m.vmEnv.Define("returnFormat", m.returnFormat)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
	m.Build(0, 0)
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&m.MouseTracker.BasicEntity, &m.MouseTracker.MouseComponent, nil, nil)
		}
	}
}

func NewMapSystem(w *ecs.World, g *graphic.GraphicsFactory, scriptpath string) *MapSystem {
	log.Println("MapSystem was added to the Scene")
	var m MapSystem
	m.World = w
	m.GraphicFactory = g
	m.MouseTracker.BasicEntity = ecs.NewBasic()
	m.MouseTracker.MouseComponent = common.MouseComponent{Track: true}
	m.scriptpath = scriptpath
	tmpscript, err := ioutil.ReadFile(m.scriptpath)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
	m.script = string(tmpscript)
	m.vmEnv = vm.NewEnv()
	err = m.vmEnv.Define("println", fmt.Println)
	err = m.vmEnv.Define("returnFormat", m.returnFormat)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
	m.Build(0, 0)
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&m.MouseTracker.BasicEntity, &m.MouseTracker.MouseComponent, nil, nil)
		}
	}
	return &m
}
