package main

import (
	"fmt"
)

type I1 interface {
	Foo() int
}

type S1 struct {
	Value int
}

type S2 struct {
	Value int
}

func (s S1) Foo() int {
	return s.Value
}

func main() {
	i1, ok := any(S1{}).(I1)
	fmt.Println(i1, ok)

	i2, ok := any(S2{}).(I1)
	fmt.Println(i2, ok)
}

