package main

import (
	"fmt"
)

type I interface {
	Foo()
}

type S1 struct {}
type S2 struct {}

func (s *S1) Foo() {}
func (s *S2) Foo() {}

func main() {
	var s1 I = &S1{}
	var s2 I = &S2{}
	
	fmt.Println(s1 == s2, s1 == s1)
}
