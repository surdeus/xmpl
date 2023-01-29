package main

import (
	"fmt"
)

type I[T any] interface {
	GetValue() T
}

type SI I[string]

type S string

func (s S) GetValue() string {
	return string(s)
}

func main() {
	var i I[string] = S("suck")
	fmt.Println(i.GetValue())
}
