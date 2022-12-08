package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8092")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("start listening")
	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("new connection:" + tcpConn.RemoteAddr().String())
		go handle(tcpConn)
	}
}

func handle(conn *net.TCPConn) {
	conn.Close()
	return
}
