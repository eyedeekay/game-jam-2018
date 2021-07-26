package main

import (
	"bytes"
	"flag"
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/eyedeekay/game/menu"

	"golang.org/x/image/font/gofont/gosmallcaps"
)

type Game struct{}

// Type uniquely defines your game type
func (*Game) Type() string { return "Game" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*Game) Preload() {
	//	engo.Files.Load("textures/city.png")
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*Game) Setup(u engo.Updater) {
	// Basic Controls
	engo.Input.RegisterButton("RotateHUD", engo.KeyEquals)
	engo.Input.RegisterButton("Quit", engo.KeyQ)

	// Modifier keys
	engo.Input.RegisterButton("RightControl", engo.KeyRightControl)
	engo.Input.RegisterButton("LeftControl", engo.KeyLeftControl)
	engo.Input.RegisterButton("RightShift", engo.KeyRightShift)
	engo.Input.RegisterButton("LeftShift", engo.KeyLeftShift)

	// Controls
	engo.Input.RegisterButton("Left", engo.KeyArrowLeft)
	engo.Input.RegisterButton("Right", engo.KeyArrowRight)
	engo.Input.RegisterButton("Up", engo.KeyArrowUp)
	engo.Input.RegisterButton("Down", engo.KeyArrowDown)
	engo.Input.RegisterButton("Tab", engo.KeyTab)
	engo.Input.RegisterButton("Enter", engo.KeyEnter)
	engo.Input.RegisterButton("Grave", engo.KeyGrave)

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	menusystem := &menu.MenuSystem{
		TextArea: menu.NewDiv("Welcome to the game world", 10, 10),
		MainMenu: menu.NewListing(engo.WindowWidth()/3, engo.WindowHeight()/3, []string{
			"Start Game",
			"Game Options",
			"Exit",
		},
		),
		OptionsMenu: menu.NewListing(engo.WindowWidth()/3, engo.WindowHeight()/3, []string{
			"Sound",
			"Graphics",
			"Controls",
			"Language",
			"Back",
		},
		),
	}
	world.AddSystem(menusystem)
	common.SetBackground(color.RGBA{77, 77, 77, 255})
	menusystem.OptionsMenu.Hide()
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&menusystem.TextArea.BasicEntity, &menusystem.TextArea.RenderComponent, &menusystem.TextArea.SpaceComponent)
			for _, mm := range menusystem.MainMenu.Elements {
				sys.Add(&mm.BasicEntity, &mm.RenderComponent, &mm.SpaceComponent)
			}
			for _, mm := range menusystem.OptionsMenu.Elements {
				sys.Add(&mm.BasicEntity, &mm.RenderComponent, &mm.SpaceComponent)
			}
		case *menu.MenuSystem:
		default:
			log.Println("No handle defined for type:", sys)
		}
	}
}

func main() {
	flag.Parse()
	opts := engo.RunOptions{
		Title:  "mmo.i2p",
		Width:  *width,
		Height: *height,
	}
	engo.Run(opts, &Game{})
}
