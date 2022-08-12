package main

import (
	"fmt"
)

type MyType int

func Handle(v any) {
	switch v.(type) {
	case string :
		fmt.Println("string")
	case int :
		fmt.Println("int")
	case MyType :
		fmt.Println("MyType")
	default :
		fmt.Println("default (unknown type)")
	}
}

func main() {
	var (
		i int
		s string
		mt MyType
	)

	Handle(i)
	Handle(s)
	Handle(mt)
	Handle(3.5)
}
