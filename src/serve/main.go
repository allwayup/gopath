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

func closeConn(conn net.Conn) {
	if conn != nil {
		clientMap[conn.RemoteAddr().String()] = nil
		conn.Close()
	}
}

func read(conn net.Conn) {
	defer closeConn(conn)
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

func writeIndex(conn net.Conn) {
	s := []string{
		"HTTP/1.1 200 OK\r\n",
		"Content-Type:application/json\r\n",
		"Content-Length:20\r\n\r\n",
		`{"name":"zs","age":18}`,
	}
	var p string
	for _, v := range s {
		p = p + v
	}
	conn.Write([]byte(p))
}

func write(conn net.Conn) {
	defer closeConn(conn)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n')
		t := strings.Trim(s, "\r\n")
		if "q" == t {
			return
		} else if "w" == t {
			writeIndex(conn)
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
