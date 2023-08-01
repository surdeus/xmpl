package main

import (
	"fmt"
	"errors"
)

type Int int

type Person struct {
	Name string
        Surname string
}

func (p Person) String() string {
	return fmt.Sprintf("%s\n%s", p.Name, p.Surname)
}

func (i Int) Inc() Int {
	i++
	return i
}

func Foo(i int) (int, error) {
	if i > 0 {
		return 0, errors.New("FUCK YOU")
	}

	return -i, nil
}