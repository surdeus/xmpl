package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, "Hello, World!\nHello, Cock!")
}

func (g *Game) Layout(ow, oh int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	err := ebiten.RunGame(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
