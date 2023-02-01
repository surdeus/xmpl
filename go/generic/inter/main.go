package main

import (
	"fmt"
)

type I[T any] interface {
	GetValue() T
}

type SI I[string]

type S string

func (s S) GetValue() SI {
	return s
}

func main() {
	var i I[SI] = S("suck")
	fmt.Println(i.GetValue())
}
