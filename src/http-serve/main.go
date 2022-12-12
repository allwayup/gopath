package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
	for {
		if conn == nil {
			return
		}
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("conn.Read():", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("客户端发来的: ", recvStr)
	}
}

func writeIndex(conn net.Conn) {
	s := []string{
		"HTTP/1.1 200 OK\r\n",
		"Content-Type:application/json;charset=UTF-8\r\n",
		"Content-Length:58\r\n\r\n",
		`{"code":200,"msg":"success","data":{"name":"zs","age":18}}`,
	}
	for _, v := range s {
		conn.Write([]byte(v))
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
		log.Println("net.Listen():", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listen.Accept():", err)
			continue
		}
		key := conn.RemoteAddr().String()
		saveConn := clientMap[key]
		if saveConn == nil {
			fmt.Println("client of:", key)
			clientMap[key] = conn
			go writeIndex(conn)
			go read(conn)
		}
	}
}
