package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Q struct {
	X, Y int
	S string
}

func main() {
	var (
		err error
		buf bytes.Buffer
	)

	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err = enc.Encode(Q{1, 2, "string value"})
	if err != nil {
		log.Fatal(err)
	}

	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", q)
}
	
