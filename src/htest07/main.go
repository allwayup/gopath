package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s := fmt.Sprintf("%s", "\u4DC0")
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(hex.EncodeToString(bs))
	// e4b780
	// 1110 0100 1011 0111 1000 0000
	// 0100 1101 1100 0000
}
