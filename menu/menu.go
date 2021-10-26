package menu

import (
	"log"
	"os"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type MenuSystem struct {
	TextArea *Div
	World    *ecs.World
	Tracker  *MouseTracker

	MainMenu    *Listing
	OptionsMenu *Listing

	MenuMax   int
	MenuIndex int

	MenuControl bool
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (m *MenuSystem) Remove(be ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (m *MenuSystem) Update(dt float32) {
	//
	if engo.Input.Button("RightShift").Down() || engo.Input.Button("LeftShift").Down() {
		if engo.Input.Button("RotateHUD").JustPressed() {
			m.TextArea.ExpandHUD()
			log.Println("Expanding HUD", m.TextArea.Corner)
		}
	} else {
		if engo.Input.Button("RotateHUD").JustPressed() {
			m.TextArea.SwitchCorner()
			log.Println("Rotating HUD", m.TextArea.Corner)
		}
	}

	//
	if engo.Input.Button("RightControl").Down() || engo.Input.Button("LeftControl").Down() {
		if engo.Input.Button("Quit").JustPressed() {
			os.Exit(0)
		}
	} else {
	}

	if engo.Input.Button("Enter").JustPressed() {
		if m.MenuControl {
			switch m.MenuIndex {
			case 1:
				switch m.OptionsMenu.index {
				case 1:
				case 2:
				case 3:
				case 4:
					m.OptionsMenu.Hide()
					m.MainMenu.Show()
					m.MenuIndex = 0
				default:

				}

			default:
				switch m.MainMenu.index {
				case 1:
					m.MainMenu.Hide()
					m.OptionsMenu.Show()
					m.MenuIndex = 1
				case 2:
					os.Exit(0)
				default:
					m.MainMenu.Hide()
					m.MenuControl = false
				}
			}
		}
	}
	if engo.Input.Button("Grave").JustPressed() {
		if !m.MenuControl {
			m.MainMenu.Show()
			m.MenuControl = true
		} else {
			m.MainMenu.Hide()
			m.OptionsMenu.Hide()
			m.MenuControl = false
			m.MenuIndex = 0
		}
	}

	if m.MenuControl {
		if engo.Input.Button("Right").JustPressed() {

			switch m.MenuIndex {
			case 1:
				m.OptionsMenu.SwitchIndexDown()
			default:
				m.MainMenu.SwitchIndexDown()
			}
		}
		if engo.Input.Button("Left").JustPressed() {
			switch m.MenuIndex {
			case 1:
				m.OptionsMenu.SwitchIndexUp()
			default:
				m.MainMenu.SwitchIndexUp()
			}
		}
		if engo.Input.Button("Down").JustPressed() {
			switch m.MenuIndex {
			case 1:
				m.OptionsMenu.SwitchIndexUp()
			default:
				m.MainMenu.SwitchIndexUp()
			}
		}
		if engo.Input.Button("Up").JustPressed() {
			switch m.MenuIndex {
			case 1:
				m.OptionsMenu.SwitchIndexDown()
			default:
				m.MainMenu.SwitchIndexDown()
			}
		}
	} else {
		//At this point, menu control swtich is disabled and everything
		//is a signed message to an entity under our control who in turn
		//will send a message to every peer who is maintaining the state
		//of the game.
		if engo.Input.Button("Right").Down() {
		}
		if engo.Input.Button("Left").Down() {
		}
		if engo.Input.Button("Down").Down() {
		}
		if engo.Input.Button("Up").Down() {
		}
	}

}

// New is the initialisation of the System
func (m *MenuSystem) New(e *ecs.World) {
	m.World = e
	m.MenuControl = true

	m.Tracker = &MouseTracker{}
	m.Tracker.BasicEntity = ecs.NewBasic()
	m.Tracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range m.World.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&m.Tracker.BasicEntity, &m.Tracker.MouseComponent, nil, nil)
		}
	}

	log.Println("MenuSystem was added to the Scene")
}
