package main

import (
	"fmt"
)

type Parent struct {
	x, y int
}

type Child struct {
	Parent
}

func (p Parent)Foo() {
	fmt.Println(p.x, p.y)
}

func NewChild(x, y int) Child {
	return Child{Parent{x, y}}
}

func main() {
	p := Parent{1, 2}
	p.Foo()

	c := NewChild(3, 4)
	c.Foo()
}
