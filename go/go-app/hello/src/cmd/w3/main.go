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
	name string
	names []string
}

type Parent struct {
	app.Compo
}

type NameList struct {
	app.Compo
	Names []string
}

func (n *NameList) Render() app.UI {
	names := n.Names
	return app.Div().Class("name-list").
		Body(
			app.H2().Class("title").Text("Names"),
			app.Range(names).Slice(func(i int) app.UI{
				return app.Div().
					Class("name").
					Text(names[i])
			}),
		)
}

func (n *NameList)OnUpdate(ctx app.Context) {
}

func (n *NameList)SetNames(names []string) app.UI {
	n.Names = names
	return n
}

func (n *NameList)Add(name string) {
	n.Names = append(n.Names, name)
}


func (h *Hello) Render() app.UI {
	inputVal := ""

	label := app.Label().
		Text("Name")

	placeholder := app.Div().Class("placeholder").
		Text("length in between of 5 and 10 chars")

	input := app.Input().
		Type("text").
		Required(true).
		Attr("minlength", 5).
		Attr("maxlength", 10).
		Value(inputVal).
		OnChange(h.ValueTo(&h.name)).
		AutoFocus(true)

	button := app.Input().
		Type("button").
		OnClick(func(ctx app.Context, e app.Event){
			h.names = append(h.names, h.name)
		}).
		Value("Submit")

	return  app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "",
			app.Text(h.name),
			).Else(
				app.Text("World"),
			),
			app.Text("!"),
		),
		app.Div().Class("field name").Body(
			label,
			placeholder,
			input,
			button,
		),
		(&NameList{Names: h.names}),
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
		Styles: []string {
			"/s/style/hello.css",
		},
	}

	s := r.AppWASM()
	fmt.Printf("'%s', '%s', '%s'\n", r.Package(), r.Static(), s) 

	fs := http.FileServer(http.Dir("dat/s"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))
	http.Handle("/", h)

	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal(err)
	}
}

