package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	ebiten_menu "github.com/hinst/ebiten-menu"
)

type Game struct {
	menu ebiten_menu.MenuUserInterface
}

func (me *Game) Initialize() {
	me.menu = ebiten_menu.MenuUserInterface{
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
	return nil
}

func (me *Game) Draw(screen *ebiten.Image) {
	me.menu.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
