package main

import (
	"fmt"
)

func main() {
	//var value = 10
	//var value1, value2 int
	if 1 < 2 {
		fmt.Println("Hello, World!")
	} else {
		fmt.Println("Fuck you!")
	}

	value := float32(10000)
	value /= 4
	fmt.Println(value)

	array := [32]int{}
	array[6] = 56
	fmt.Println(array)

	slice := []int{234, 2321, 5465, 34}
	slice = append(slice, 32)
	fmt.Println(slice)

	// "Range" loop.
	for i, v := range slice {
		fmt.Println(i, v)
	}

	if _, err := Foo(2) ; err != nil {
		panic(err)
	}

	// "While" loop.
	i := 0
	for i<10 {
		fmt.Println(i)
		i++
	}
	for i := range slice{
		slice[i]*=13
		
	}

	fmt.Println(slice)
	// "For" loop.
	for i := i ; i < 10 ; i++ {
		fmt.Println(i)
	}

	p1 := Person{
		Name: "Alexander",
		Surname: "Gulyaev",
	}

	fmt.Println(p1)

	myInt := Int(1)
	fmt.Println(myInt, myInt.Inc(), myInt)
}

func sum(a, b int) int {
	return a + b
}