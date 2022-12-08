package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	clientMap = make(map[string]net.Conn)
)

func read(conn net.Conn) {
	if conn == nil {
		fmt.Println("read close conn")
		return
	}
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			return
		}
		fmt.Printf("%s send: %s\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
}

func write(conn net.Conn) {
	defer (func() {
		if conn != nil {
			conn.Close()
		}
	})()

	keyIn := bufio.NewReader(os.Stdin)
	for {
		if conn == nil {
			fmt.Println("write close conn")
			return
		}

		readString, err := keyIn.ReadString('\n')
		if err != nil {
			continue
		}
		trim := strings.Trim(readString, "\r\n")
		if strings.ToUpper(trim) == "Q" {
			return
		}
		_, err = conn.Write([]byte(trim))
		if err != nil {
			continue
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go read(conn)
		key := conn.RemoteAddr().String()
		saveConn := clientMap[key]
		if saveConn == nil {
			clientMap[key] = conn
			go write(conn)
		}
	}
}
