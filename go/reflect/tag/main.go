package main

import (
	"reflect"
	"fmt"
)

func main() {
	anon := struct{
		Value1 int `op:"shit"`
		String1 string `op:"other shit"`
	}{
		228,
		"Hello, World!",
	}

	t := reflect.TypeOf(anon)
	fmt.Println(
		t.Field(0).Tag.Get("op"),
		t.Field(1).Tag.Get("op"),
	)
}
