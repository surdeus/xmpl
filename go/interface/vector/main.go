package main

import(
	"fmt"
	"vector/m/vector1"
	"vector/m/vector2"
)

func main() {
	v1 := vector1.Vector{1, 1}
	v2 := vector2.Vector{2, 2}
	fmt.Println(vector1.Add(v1, v2))
}
