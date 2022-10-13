package main

import (
	"html/template"
	"fmt"
	"os"
)

func main() {
	//tmpl := template.Must(template.ParseFiles("dir/template"))
	b, _ := os.ReadFile(string("dir/template"))
	tmpl := template.New("dir/template")
	tmpl, _ = tmpl.Parse(string(b))
	tmpl = tmpl.New("other name")
	tmpl, _ = tmpl.Parse("other template")
	tmpl = tmpl.New("2")
	tmpl, _ = tmpl.Parse("2 template")

	tmpl.ExecuteTemplate(os.Stdout, "dir/template", nil)

	fmt.Println("")

	tmpl.ExecuteTemplate(os.Stdout, "other name", nil)

	fmt.Println("")
	tmpl.ExecuteTemplate(os.Stdout, "2", nil)
}
