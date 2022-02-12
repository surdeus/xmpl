package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/colornames"
	"fmt"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(500, 500), basicAtlas)

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicTxt.Color = colornames.Green
	fmt.Fprintln(basicTxt, "Hello, Text!")

	basicTxt.Color = colornames.Blue
	fmt.Fprintln(basicTxt, "I support multiple lines!")

	basicTxt.Color = colornames.Red
	fmt.Fprintln(basicTxt, "And I'm an %s, yay!", "io.Writer")

	for !win.Closed() {
		win.Clear(colornames.Black)
		basicTxt.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

