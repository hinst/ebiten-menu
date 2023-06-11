package ebiten_menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuUserInterface struct {
	Items []MenuUserInterfaceItem
}

type MenuUserInterfaceItem struct {
	Title string
	Id    int
}

func (me *MenuUserInterface) Update(deltaTime float64) {
}

func (me *MenuUserInterface) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Good day", 1, 1)
}
