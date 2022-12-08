package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net"
)

type SyncVotePool struct {
	Id     string
	Name   string
	Ip     string
	Port   string
	Serves []*SyncVotePool
}

func connect() {
	conn, err := net.Dial("tcp", "43.250.33.235:443")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("client success")
	var buf bytes.Buffer
	_, err = io.Copy(&buf, conn)
	if err != nil {
		panic(err)
	}
	log.Println(buf)
}

func main() {
	var v SyncVotePool = SyncVotePool{
		Id:     "1",
		Name:   "zs",
		Serves: []*SyncVotePool{&SyncVotePool{Id: "2"}},
	}
	s, _ := json.Marshal(v)
	log.Println(string(s))
}
