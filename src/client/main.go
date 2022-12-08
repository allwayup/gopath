package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	Protocol = "tcp"
	Addr     = "127.0.0.1:20000"
)

var (
	inputReader *bufio.Reader
)

func read(conn net.Conn) {
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("client read: ", err)
		os.Exit(0)
		return
	}
	fmt.Println("服务端发来的: ", string(buf[:n]))
}

func write(conn net.Conn) {
	go read(conn)
	input, _ := inputReader.ReadString('\n')
	inputInfo := strings.Trim(input, "\r\n")
	if strings.ToUpper(inputInfo) == "Q" {
		return
	}
	_, err := conn.Write([]byte(inputInfo))
	if err != nil {
		fmt.Println("client write: ", err)
	}
}

func main() {
	conn, err := net.Dial(Protocol, Addr)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer conn.Close()

	inputReader = bufio.NewReader(os.Stdin)
	for {
		write(conn)
	}
}
