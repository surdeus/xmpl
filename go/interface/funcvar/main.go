package main

import (
	"fmt"
)

type S struct {
	Foo func(a, b int) int
}

type I interface {
	Foo(a, b int) int
}

func main() {
	s := S{}
	i := I(s)
	fmt.Println(i)
}

