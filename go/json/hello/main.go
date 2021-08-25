package main

import(
	"encoding/json"
	"fmt"
	"log"
)


type User struct {
	Id int
	Name, Occupation string
}

func fatalErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var err error

	// Encoding.
	u1 := User{1, "John Doe", "gardener"}
	d1, err := json.Marshal(u1)
	fatalErrorCheck(err)

	fmt.Println(string(d1))

	// Decoding.
	d2 := []byte(`{
		"Id":2,
		"Name":"Donald Dickson",
		"Occupation":"striper"}`)
	var u2 User
	err = json.Unmarshal(d2, &u2)
	fatalErrorCheck(err)

	// Encoding again for output.
	d22, err := json.Marshal(u2)
	fatalErrorCheck(err)
	fmt.Println(string(d22))
}
