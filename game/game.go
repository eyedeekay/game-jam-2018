package gamemap

import (
	//	"log"
	"image/color"
	"path/filepath"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"../config"
	//entity "../entity"
	graphic "../graphic"
	"../system"
)

type WholeGame struct {
	title      string
	assets     string
	configfile string
	skel       string

	config         *config.Conf
	graphicFactory *graphic.GraphicsFactory

	headlessMode   bool
	standardInputs bool
	fullscreen     bool
	vsync          bool
	resizable      bool

	fps    int
	width  int
	height int
}

func (w WholeGame) Type() string { return "gamewrapper" }

func (w WholeGame) Preload() {
	engo.Files.Load(filepath.Join("icon", "icon.png"))
}

func (w WholeGame) Setup(u engo.Updater) {
	common.SetBackground(color.NRGBA{R: 125, G: 125, B: 125, A: 125})
	engo.Input.RegisterButton("AddWall", engo.KeyF1)
	engo.Input.RegisterButton("Exit", engo.KeyEscape)
	engo.Input.RegisterButton("Left", engo.KeyA)
	engo.Input.RegisterButton("Right", engo.KeyD)
	engo.Input.RegisterButton("Up", engo.KeyW)
	engo.Input.RegisterButton("Down", engo.KeyS)

	world, _ := u.(*ecs.World)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&common.MouseZoomer{0.125})

	world.AddSystem(common.NewKeyboardScroller(
		400, engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis))
	world.AddSystem(&common.EdgeScroller{400, 20})

	world.AddSystem(systems.NewMapSystem(world, w.graphicFactory, "assets/scripts/map/demo.ank"))
	world.AddSystem(&systems.ScaleSystem{})
}

func (w WholeGame) Options() engo.RunOptions {
	return engo.RunOptions{
		Title:          w.title,
		Width:          w.width,
		Height:         w.height,
		HeadlessMode:   w.headlessMode,
		Fullscreen:     w.fullscreen,
		StandardInputs: w.standardInputs,
		FPSLimit:       w.fps,
		AssetsRoot:     w.assets,
		VSync:          w.vsync,
		NotResizable:   w.resizable,
	}
}

func NewGame(opts ...func(*WholeGame) error) (*WholeGame, error) {
	var w WholeGame
	w.title = "gamewrapper"
	w.assets = "assets"
	w.configfile = "config.ini"
	w.width = 640
	w.height = 640
	w.fps = 40
	w.fullscreen = false
	w.headlessMode = false
	w.standardInputs = true
	w.skel = "skel"
	w.vsync = false
	w.resizable = false

	var err error
	for _, o := range opts {
		if err := o(&w); err != nil {
			return &w, err
		}
	}
	common.CameraBounds = engo.AABB{Min: engo.Point{0, 0}, Max: engo.Point{800, 800}}
	common.MinZoom = .25
	common.MaxZoom = 3
	w.config, err = config.NewWholeGameConf(filepath.Join(w.assets, w.configfile))
	if err != nil {
		return nil, err
	}
	colorset, b := w.config.Get("colorset")
	if !b {
		colorset = "colors.txt"
	}
	if err != nil {
		return nil, err
	}
	w.graphicFactory, err = graphic.NewGraphicsFactory(w.assets, colorset, w.skel, w.configfile)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (w WholeGame) RunScene() {
	engo.Run(w.Options(), w)
}
