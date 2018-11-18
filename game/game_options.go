package gamemap

type Option func(*WholeGame) error

//SetTitle sets the host of the client's SAM bridge
func SetTitle(s string) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.title = s
		return nil
	}
}

//SetWidth sets the width of the window
func SetWidth(s int) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.width = s
		return nil
	}
}

//SetHeight sets the height of the window
func SetHeight(s int) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.height = s
		return nil
	}
}

//SetHeadless sets the game to server-only mode
func SetHeadless(s bool) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.headlessMode = s
		return nil
	}
}

//SetFullscreen sets the game to full-screen mode
func SetFullscreen(s bool) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.fullscreen = s
		return nil
	}
}

//SetVsync sets the game to full-screen mode
func SetVsync(s bool) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.vsync = s
		return nil
	}
}

//SetResizable sets the game to full-screen mode
func SetResizable(s bool) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.resizable = s
		return nil
	}
}

//SetTitle sets the host of the client's SAM bridge
func SetStandardInputs(s bool) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.standardInputs = s
		return nil
	}
}

//SetFPS sets the frame-per-second limit
func SetFPS(s int) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.fps = s
		return nil
	}
}

//SetAssets sets the assets folder.
func SetAssets(s string) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.assets = s
		return nil
	}
}

//SetConfigFile sets the filename of the config file in the assets folder
func SetConfigFile(s string) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.configfile = s
		return nil
	}
}

//SetSekFolder sets the filename of the config file in the assets folder
func SetSkelFolder(s string) func(*WholeGame) error {
	return func(w *WholeGame) error {
		w.skel = s
		return nil
	}
}
