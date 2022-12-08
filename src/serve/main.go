package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LogPath  = "./log.txt"
	Protocol = "tcp"
	Addr     = "127.0.0.1:20000"
)

var (
	clientMap = make(map[string]net.Conn)
)

func read(conn net.Conn) {
	defer (func() {
		if conn != nil {
			defer conn.Close()
		}
	})()
	for {
		if conn == nil {
			return
		}
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("客户端发来的: ", recvStr)
	}
}

func write(conn net.Conn) {
	defer (func() {
		if conn != nil {
			defer conn.Close()
		}
	})()
	for {
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n')
		t := strings.Trim(s, "\r\n")
		if "Q" == t {
			return
		}
		conn.Write([]byte(t))
	}
}

func main() {
	f, err := os.Open(LogPath)
	if err != nil {
		f, err = os.Create(LogPath)
		if err != nil {
			panic(err)
		}
	}
	log.SetOutput(f)

	listen, err := net.Listen(Protocol, Addr)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		key := conn.RemoteAddr().String()
		saveConn := clientMap[key]
		if saveConn == nil {
			fmt.Println("client of:", key)
			clientMap[key] = conn
			go write(conn)
			go read(conn)
		}
	}
}
