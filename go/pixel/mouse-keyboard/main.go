package main

import (
	"image"
	"os"
	"time"
	"math"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	//"fmt"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture("../hiking.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())


	movSpeed := 200.0
	pos := win.Bounds().Center()
	mat := pixel.IM
	movmat := pixel.IM
	rotmat := pixel.IM

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Whitesmoke)

		if win.Pressed(pixelgl.MouseButtonLeft){
			click := win.MousePosition()
			direction := click.Sub(pos) 
			angle := math.Atan(direction.Y/direction.X)
			if direction.X < 0 { angle += math.Pi }
			rotmat = pixel.IM.Rotated(pixel.ZV, angle)
		}
		if win.Pressed(pixelgl.KeyUp){
			pos.Y += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyDown){
			pos.Y -= movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyRight){
			pos.X += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyLeft){
			pos.X -= movSpeed*dt
		}

		movmat = pixel.IM.Moved(pos)
		mat = rotmat.Chained(movmat)
		sprite.Draw(win, mat)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

