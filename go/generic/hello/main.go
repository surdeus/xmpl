package main

/* Go demonstration of generics. */

import (
	"fmt"
)

func Sum[V int | int64 | float64 | string](a []V) V {
	var s V
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	fmt.Println(Sum[int]([]int{1, 2, 3, 4, 10}))
}

