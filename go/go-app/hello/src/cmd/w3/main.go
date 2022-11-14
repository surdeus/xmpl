package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"net/http"
	"log"
	"fmt"
)

type Res struct {}
func (r *Res)Package() string { return "" }
func (r *Res)Static() string { return "/s" }
func (r *Res)AppWASM() string { return "/s/w3.wasm"}

type Hello struct {
	app.Compo
	Name string
}

func (h *Hello) Render() app.UI {
	fmt.Println("yes")
	return app.H1().Body(
		app.Strong().Text("Hello, New App!"),
		app.Div().Text("Another fuck!!!"),
	)
}

func main() {
	app.Route("/", &Hello{})
	app.RunWhenOnBrowser()

	r := &Res{}
	h := &app.Handler {
		Name: "Hello",
		Description: "dick",
		Resources: r,
	}


	s := r.AppWASM()
	fmt.Printf("'%s', '%s', '%s'\n", r.Package(), r.Static(), s) 

	fs := http.FileServer(http.Dir("dat/s"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))
	http.Handle("/", h)

	fmt.Println("Running on ''")
	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal(err)
	}
}

