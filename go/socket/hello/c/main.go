package main

import(
	"fmt"
	"net"
	"os"
	"bufio"
)

func
main(){
	c, _ := net.Dial("tcp", "127.0.0.1:8081")
	for{
		r := bufio.NewReader(os.Stdin)
		cr := bufio.NewReader(c)
		txt, _ := r.ReadString('\n')
		fmt.Fprint(c, txt + "\n")
		msg, _ := cr.ReadString('\n')
		fmt.Print(msg)
	}
}
