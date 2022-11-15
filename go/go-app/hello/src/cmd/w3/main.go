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

type Root struct {
	app.Compo
	Page string
	Pages map[string] app.UI
}

type Hello struct {
	app.Compo
	nameList *NameList
	greeter *Greeter
	name string
}

type Parent struct {
	app.Compo
}

type NameList struct {
	app.Compo
	Names []string
}

type Greeter struct {
	app.Compo
	Name string
}

func (g *Greeter) Render() app.UI {
	return app.H1().Body(
			app.Text("Hello, "),
			app.If(g.Name != "",
			app.Text(g.Name),
			).Else(
				app.Text("World"),
			),
			app.Text("!"),
	)
}

func (p *Parent)Render() app.UI {
	return app.Div().Class("parent").Body(
		app.H1().Text("The other shit"),
	)
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

func (n *NameList)Add(name string) {
	n.Names = append(n.Names, name)
	n.Update()
}

func (g *Greeter)SetName(name string) {
	g.Name = name
	g.Update()
}

func (h *Hello) Render() app.UI {
	var (
		input app.UI
	)
	label := app.Label().
		Text("Name")

	placeholder := app.Div().Class("placeholder").
		Text("length in between of 5 and 10 chars")

	submitFunc := func(ctx app.Context, e app.Event){
		h.nameList.Add(h.greeter.Name)
		h.greeter.SetName("")
		input.JSValue().Call("focus")
	}

	input = app.Input().Class("1").Class("2").
		Type("text").
		Required(true).
		Attr("minlength", 5).
		Attr("maxlength", 10).
		Value(h.greeter.Name).
		OnInput(func(ctx app.Context, e app.Event) {
			v := ctx.JSSrc().Get("value")
			h.greeter.SetName(v.String())
		}).
		OnKeyPress(func(ctx app.Context, e app.Event) {
			if e.Value.Get("key").String() == "Enter" {
				submitFunc(ctx, e)
			}
		}).
		AutoFocus(true)

	button := app.Input().
		Type("button").
		OnClick(submitFunc).
		Value("Submit")

	return app.Div().Class("hello").Body(
		h.greeter,
		app.Div().Class("field name").Body(
			label,
			placeholder,
			input,
			button,
		),
		h.nameList,
	)
}

func (r *Root) OnMount(ctxt app.Context) {
	r.Pages = map[string] app.UI {
		"hello" : &Hello{nameList: &NameList{}, greeter: &Greeter{} },
		"parent" : &Parent{},
	}
	r.Page = "hello"
}

func (h *Hello) OnMount(ctx app.Context) {
	//h.nameList = &NameList{Names: []string{"cocks", "sucks"}}
	//h.greeter = &Greeter{Name: ""}
}

func (r *Root)Render() app.UI {
	nav := app.Nav().Class("nav").Body(
		app.Input().
			Type("button").
			Value("hello").
			OnClick( func(c app.Context, e app.Event) {
				r.Page = "hello"
			}),
		app.Input().
			Type("button").
			Value("parent").
			OnClick( func(c app.Context, e app.Event) {
				r.Page = "parent"
			}),

	)

	page, _ := r.Pages[r.Page]
	return app.Div().Class("root").Body(
		nav,
		page,
	)
}


func main() {
	app.Route("/", &Root{} )
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

