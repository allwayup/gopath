package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("./log.txt")
	if err != nil {
		panic(err)
	
	log.SetOutput(f)
	log.Println("------")
}
