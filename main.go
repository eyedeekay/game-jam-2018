package main

import (
	"flag"
	"log"
	"math/rand"

	//ent "./entity"
	game "./game"
)

func RandWord() string {
	b := make([]byte, 12)
	for i := range b {
		b[i] = "abcdefghijklmnopqrstuvwxyz"[rand.Intn(len("abcdefghijklmnopqrstuvwxyz"))]
	}
	return string(b)
}

var (
	title        = flag.String("title", RandWord(), "name to give the game server")
	config       = flag.String("config", "config.ini", "config file name(in assets/ folder)")
	assets       = flag.String("assets", "assets", "assets root")
	width        = flag.Int("width", 1280, "width of the window")
	height       = flag.Int("height", 736, "height of the window")
	fps          = flag.Int("fps", 40, "limit FPS")
	headless     = flag.Bool("headless", false, "run windowless")
	fullscreen   = flag.Bool("fullscreen", false, "run fullscreen")
	inputs       = flag.Bool("inputs", true, "use default inputs")
	vsync        = flag.Bool("vsync", false, "enable vsync")
	notresizable = flag.Bool("noresize", true, "enable resizable window")
)

func main() {
	flag.Parse()
	if g, e := game.NewGame(
		game.SetTitle(*title),
		game.SetAssets(*assets),
		game.SetWidth(*width),
		game.SetHeight(*height),
		game.SetHeadless(*headless),
		game.SetFullscreen(*fullscreen),
		game.SetStandardInputs(*inputs),
		game.SetFPS(*fps),
		game.SetConfigFile(*config),
		game.SetVsync(*vsync),
		game.SetResizable(*notresizable),
	); e == nil {
		g.RunScene()
	} else {
		log.Printf("Main Loading Error: %s", e.Error())
	}
}
