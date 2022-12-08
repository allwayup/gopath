package main

import (
	"fmt"
	"reflect"
)

func main() {
	//s := "`1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM[];',./~!@#$%^&*()_+-={}:\"<>?"
	s := "1.统计单字符重复个数;2.并排序"
	r := []rune(s)

	h := make(map[int32]int32)
	l := len(r)
	for i := 0; i < l; i++ {
		v := r[i]
		h[v] = v
		fmt.Println(reflect.TypeOf(v))
	}

	fmt.Println("end.............")
}
