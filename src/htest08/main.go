package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		c, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		var b []byte
		c.Read(b)
		fmt.Println(string(b))
	}
}
