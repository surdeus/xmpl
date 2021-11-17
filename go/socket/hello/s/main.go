package main

import(
	"net"
	"bufio"	
	"strings"
)

func
main(){
	l, _ := net.Listen("tcp", ":8081")
	c, _ := l.Accept()
	for{
		m, _ := bufio.NewReader(c).ReadString('\n')
		c.Write([]byte(strings.ToUpper(m) + "\n"))
	}
}
