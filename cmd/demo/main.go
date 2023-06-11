package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebiten_menu "github.com/hinst/ebiten-menu"
)

type Game struct {
	UpdateTime      time.Time
	JustPressedKeys []ebiten.Key
	Menu            ebiten_menu.MenuUserInterface
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
				Title: "Information",
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
}

func (me *Game) draw(screen *ebiten.Image) {
	me.Menu.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	var game = &Game{}
	game.Initialize()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
