package main

import (
	"fmt"
)
type Ret struct{I int ; V string}

func Foo() chan Ret {
	ret := make(chan Ret)

	go func() {
		ret <- Ret{1, "die"}
		ret <- Ret{5, "cock"}
		ret <- Ret{8, "suck"}
		ret <- Ret{7, "suck"}
		close(ret)
	}()

	return ret
}

func main() {
	for c := range Foo() {
		fmt.Println(c.I, c.V)
	}
}
