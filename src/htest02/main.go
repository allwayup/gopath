package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handler(conn net.Conn) {
	keyIn := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(conn)
	if conn == nil {
		return
	}

	readString, serr := keyIn.ReadString('\n')
	if serr != nil {
		return
	}
	trim := strings.Trim(readString, "\r\n")
	if strings.ToUpper(trim) == "Q" {
		return
	}
	_, err := conn.Write([]byte(trim))
	if err != nil {
		return
	}

	var buf [128]byte
	n, rerr := reader.Read(buf[:])
	if rerr == nil {
		fmt.Printf("%s send: %s\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
}

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8084")
	if err != nil {
		panic(err)
	}
	defer (func() {
		if dial != nil {
			dial.Close()
		}
	})()
	for {
		handler(dial)
	}
}
