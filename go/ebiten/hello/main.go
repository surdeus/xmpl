package main

import (
	x "github.com/hajimehoshi/ebiten/v2"
	ux "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(s *x.Image) {
	ux.DebugPrint(s, "Hello, World!\nHello, Cock!")
}

func (g *Game) Layout(ow, oh int) (int, int) {
	return 320, 240
}

func main() {
	x.SetWindowSize(640, 480)
	x.SetWindowTitle("Hello, World!")
	err := x.RunGame(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
