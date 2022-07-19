package main

import(
	"fmt"
)

type S struct {
	i int `default:1`
	s string `default:"default string"`
	ss []string `default:[]string{}`
}


func
main(){
	s := S{}
	fmt.Println(s.i, s.s, s.ss)
}
