package ebiten_menu

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MenuUserInterface struct {
	PulseAnimationCounter float64
	Items                 []MenuUserInterfaceItem
	SelectedItemIndex     int
}

type MenuUserInterfaceItem struct {
	Title string
	Id    int
}

func (me *MenuUserInterface) Update(deltaTime float64, justPressedKeys []ebiten.Key) {
	if len(me.Items) > 0 {
		for _, key := range justPressedKeys {
			if key == ebiten.KeyUp {
				me.SelectedItemIndex -= 1
				if me.SelectedItemIndex < 0 {
					me.SelectedItemIndex = 0
				}
			}
			if key == ebiten.KeyDown {
				me.SelectedItemIndex += 1
				if len(me.Items) <= me.SelectedItemIndex {
					me.SelectedItemIndex = len(me.Items) - 1
				}
			}
		}
	} else {
		me.SelectedItemIndex = -1
	}
	me.PulseAnimationCounter += deltaTime
	if me.PulseAnimationCounter >= math.Pi {
		me.PulseAnimationCounter = 0
	}
}

func (me *MenuUserInterface) Draw(screen *ebiten.Image) {
	var menuY = screen.Bounds().Dy()/2 - len(me.Items)*me.GetCharacterHeight()/2
	for index, item := range me.Items {
		ebitenutil.DebugPrintAt(screen,
			item.Title,
			me.GetLeftMargin(),
			menuY+index*me.GetCharacterHeight(),
		)
	}
	if me.SelectedItemIndex >= 0 {
		vector.StrokeRect(
			screen,
			float32(me.GetLeftMargin())-1,
			float32(menuY)+float32(me.SelectedItemIndex*me.GetCharacterHeight()),
			2+float32(me.GetCharacterWidth()*len(me.Items[me.SelectedItemIndex].Title))+2,
			float32(me.GetCharacterHeight()),
			1,
			color.RGBA{R: 0, G: 255, B: 0, A: 0},
			false,
		)
	}
}

func (me *MenuUserInterface) GetCharacterWidth() int {
	return 6
}

func (me *MenuUserInterface) GetCharacterHeight() int {
	return 16
}

func (me *MenuUserInterface) GetLeftMargin() int {
	return 16
}
