package main

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebiten_menu "github.com/hinst/ebiten-menu"
)

const (
	LAYOUT_SMALL = iota
	LAYOUT_BIG
)

type Game struct {
	UpdateTime      time.Time
	JustPressedKeys []ebiten.Key
	Menu            ebiten_menu.MenuUserInterface
	IsExiting       bool
	LayoutMode      int
}

func (me *Game) Initialize() {
	me.UpdateTime = time.Now()
	me.Menu = ebiten_menu.MenuUserInterface{
		Items: []ebiten_menu.MenuUserInterfaceItem{
			{
				Title: "New Game",
				Id:    1,
			},
			{
				Title: "Toggle Full Screen",
				Id:    2,
			},
			{
				Title: "Exit",
				Id:    3,
			},
		},
	}
}

func (me *Game) Update() error {
	me.JustPressedKeys = inpututil.AppendJustPressedKeys(me.JustPressedKeys)
	if me.IsExiting {
		return errors.New("exiting")
	}
	return nil
}

func (me *Game) Draw(screen *ebiten.Image) {
	var updateTime = time.Now()
	me.update(updateTime.Sub(me.UpdateTime).Seconds())
	me.UpdateTime = updateTime
	me.JustPressedKeys = me.JustPressedKeys[:0]
	me.draw(screen)
}

func (me *Game) update(deltaTime float64) {
	me.Menu.Update(deltaTime, me.JustPressedKeys)
	if me.Menu.PressedItemId == 1 {
		me.LayoutMode = (me.LayoutMode + 1) % 2
	} else if me.Menu.PressedItemId == 2 {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	} else if me.Menu.PressedItemId == 3 {
		me.IsExiting = true
	}
}

func (me *Game) draw(screen *ebiten.Image) {
	me.Menu.Draw(screen)
}

func (me *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if me.LayoutMode == LAYOUT_SMALL {
		return 320, 240
	} else if me.LayoutMode == LAYOUT_BIG {
		return 640, 480
	} else {
		panic("Unknown layout: " + strconv.Itoa(me.LayoutMode))
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Menu demo")
	var game = &Game{}
	game.Initialize()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
