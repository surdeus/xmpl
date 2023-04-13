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

type P1 struct {
	V string
}

type P2 struct {
	V string
}

type C1 struct {
	P1
	P2
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
	
	c1 := C1{
		P1: P1{
			"p1",
		},
		P2: P2{"p2"},
	}
	
	fmt.Println(c1.V)
}
