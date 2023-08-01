package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, Master!")
		time.Sleep(time.Second * 1)
	}
}

