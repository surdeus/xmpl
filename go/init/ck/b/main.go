package b

import(
	"prog/c"
	"fmt"
)

func
init() {
	fmt.Println("B's init")
}

func
Foo(){
	c.Foo()
}
