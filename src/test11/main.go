package main

import (
	"fmt"
)

func p(s string) string {
	return s + "LLLLLLLL"
}

var (
	F = func() func(s string) string {
		return p
	}()
)

func main() {
	fmt.Println(F("--------"))
	fmt.Println("end...........")
}
