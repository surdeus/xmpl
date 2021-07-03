package main

import (
	"container/list"
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

type Transform struct {
	P, S pixel.Vec
	R float64
}

type Behaviorer interface {
	Start()
	Update()
	GetOD() *Object
}

type Object struct {
	T Transform
	S *pixel.Sprite
}

type GopherPlayer struct {
	O Object
	MoveSpeed float64
}

var(
	dt float64
	lasttime time.Time
	win *pixelgl.Window
	goph_sprite *pixel.Sprite
	objects *list.List
	err error
)

func
setDT(){
	dt = time.Since(lasttime).Seconds()
	lasttime = time.Now()
}

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

func
AddObject(o Behaviorer) {
	objects.PushBack(o)
	o.Start()
}

func
(g *GopherPlayer)Start() {
}

func
(g *GopherPlayer)Update() {
		pos := &(g.O.T.P)
		angle := &(g.O.T.R)
		movSpeed := g.MoveSpeed
		if win.Pressed(pixelgl.MouseButtonLeft){
			click := win.MousePosition()
			direction := click.Sub(*pos) 
			*angle = math.Atan(direction.Y/direction.X)
			if direction.X < 0 { *angle += math.Pi }

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
}

func
(g *GopherPlayer)GetOD() *Object {
	return &g.O
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	pic, err := loadPicture("../hiking.png")
	if err != nil {
		panic(err)
	}

	goph_sprite = pixel.NewSprite(pic, pic.Bounds())
	objects = list.New()

	goph_player := GopherPlayer{
		O: Object {
			T: Transform {
				P: win.Bounds().Center(),
				S: pixel.Vec{3, 1},
			},
			S: goph_sprite,
		},
		MoveSpeed: 200.0,
	}
	AddObject(&goph_player)
	
	goph_player1 := GopherPlayer{
		O: Object {
			T: Transform {
				P: pixel.ZV,
				S: pixel.Vec{1, 1},
			},
			S: goph_sprite,
		},
		MoveSpeed: 100.0,
	}
	AddObject(&goph_player1)

	lasttime = time.Now()
	for !win.Closed() {
		setDT()
		win.Clear(colornames.Whitesmoke)
		for e:= objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(Behaviorer)
			o.Update()
			od := o.GetOD()
			finmat := pixel.IM.ScaledXY(pixel.ZV, od.T.S).Rotated(pixel.ZV, od.T.R).Moved(od.T.P)
			od.S.Draw(win, finmat)
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

