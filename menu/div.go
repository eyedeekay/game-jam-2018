package menu

import (
	"image/color"
	"log"
	"math"
	"strings"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Div struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	*common.Font

	Corner   int
	Expanded bool

	Yadj float32
	Xadj float32
}

func (d *Div) ExpandedXMod() float32 {
	x := 1
	if !d.Expanded {
		x = 4
	}
	return float32(x)
}

func (d *Div) ExpandedYMod() float32 {
	y := 1
	if !d.Expanded {
		y = 6
	}
	return float32(y)
}

func (d *Div) SwitchColor(fg, bg color.Color) {
	if fg != nil {
		d.Font.FG = fg
	}
	if bg != nil {
		d.Font.BG = bg
	}

}

func (d *Div) ExpandHUD() {
	sw := !d.Expanded
	d.Expanded = sw

	newtext := d.WrapText(d.RenderComponent.Drawable.(common.Text).Text)
	d.RenderComponent.Drawable = common.Text{
		Font: d.Font,
		Text: newtext,
	}

	d.SpaceComponent.Position.X = d.RealX()
	d.SpaceComponent.Position.Y = d.RealY()
}

func (d *Div) SwitchCorner() {
	if d.Corner == 4 {
		d.Corner = 1
	} else {
		d.Corner++
	}
	d.SpaceComponent.Position.X = d.RealX()
	d.SpaceComponent.Position.Y = d.RealY()
}

func (d *Div) StartX(adj float32) float32 {
	switch d.Corner {
	// Top-Left
	case 1:
		return 0 + adj
	// Top-Right
	case 2:
		return engo.WindowWidth() - adj - (engo.WindowWidth() / d.ExpandedXMod())
	// Bottom-Right
	case 3:
		return engo.WindowWidth() - adj - (engo.WindowWidth() / d.ExpandedXMod())
	// Bottom-Left
	case 4:
		return 0 + adj
	default:
		return 0 + adj
	}
}

func (d *Div) StartY(adj float32) float32 {
	switch d.Corner {
	// Top-Left
	case 1:
		return 0 + adj
	// Top-Right
	case 2:
		return 0 + adj
	// Bottom-Right
	case 3:
		return engo.WindowHeight() - adj - (engo.WindowHeight() / d.ExpandedYMod())
	// Bottom-Left
	case 4:
		return engo.WindowHeight() - adj - (engo.WindowHeight() / d.ExpandedYMod())
	default:
		return 0 + adj
	}
}

func (d *Div) RealX() float32 {
	switch d.Corner {
	// Top-Left
	case 1:
		return 0 + d.Xadj
	// Top-Right
	case 2:
		if d.ExpandedXMod() != 1 {
			return engo.WindowWidth() - d.Xadj - (engo.WindowWidth() / d.ExpandedXMod())
		}
		return engo.WindowWidth() - (engo.WindowWidth() / d.ExpandedXMod()) + d.Xadj
	// Bottom-Right
	case 3:
		if d.ExpandedXMod() != 1 {
			return engo.WindowWidth() - d.Xadj - (engo.WindowWidth() / d.ExpandedXMod())
		}
		return engo.WindowWidth() - (engo.WindowWidth() / d.ExpandedXMod()) + d.Xadj
	// Bottom-Left
	case 4:
		return 0 + d.Xadj
	default:
		return 0 + d.Xadj
	}
}

func (d *Div) RealY() float32 {
	switch d.Corner {
	// Top-Left
	case 1:
		return 0 + d.Yadj
	// Top-Right
	case 2:
		return 0 + d.Yadj
	// Bottom-Right
	case 3:
		if d.ExpandedYMod() != 1 {
			return engo.WindowHeight() - d.Yadj - (engo.WindowHeight() / d.ExpandedYMod())
		}
		return engo.WindowHeight() - (engo.WindowHeight() / d.ExpandedYMod()) + d.Yadj
	// Bottom-Left
	case 4:
		if d.ExpandedYMod() != 1 {
			return engo.WindowHeight() - d.Yadj - (engo.WindowHeight() / d.ExpandedYMod())
		}
		return engo.WindowHeight() - (engo.WindowHeight() / d.ExpandedYMod()) + d.Yadj
	default:
		return 0 + d.Yadj
	}
}

func (d *Div) wrapChars() int {
	i := int((engo.WindowWidth() / d.ExpandedXMod()) / d.ExpandedYMod())
	//	log.Println("wrap at chars", i)
	return i
}

func (d *Div) WrapText(text string) string {
	var bytes []byte
	words := strings.Split(strings.TrimRight(strings.Replace(text, "\n", " ", -1), " "), " ")
	for index, byt := range words {
		f := float64((len(bytes) + len(byt)) / d.wrapChars())
		val := int(math.Trunc(f)+1) * d.wrapChars()
		if len(bytes)+len(byt)+index >= val {
			//			log.Println("wrapping: %s", bytes)
			bytes = append(bytes, []byte("\n")...)
		}
		bytes = append(bytes, byt...)
		bytes = append(bytes, []byte(" ")...)
	}
	//	log.Println(string(bytes))
	return strings.Replace(string(bytes), "  ", " ", -1)
}

func NewDiv(text string, xadj, yadj float32) *Div {
	return NewColorDiv(text, 3, 14, xadj, yadj, color.White, color.Black)
}

func NewColorDiv(text string, corner int, size float64, xadj, yadj float32, fg, bg color.Color) *Div {
	div := Div{BasicEntity: ecs.NewBasic()}
	div.Font = &common.Font{
		URL:  "go.ttf",
		FG:   fg,
		BG:   bg,
		Size: size,
	}
	div.Yadj = yadj
	div.Xadj = xadj
	div.Font.CreatePreloaded()
	div.RenderComponent.Drawable = common.Text{
		Font: div.Font,
		Text: div.WrapText(text),
	}
	div.SetShader(common.TextHUDShader)
	div.RenderComponent.SetZIndex(1001)
	div.Corner = corner
	div.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: div.StartX(div.Xadj), Y: div.StartY(div.Yadj)},
		Width:    engo.WindowWidth() / 2,
		Height:   engo.WindowHeight() / 3,
	}
	log.Println("Created div containing Text:", text)
	return &div
}
