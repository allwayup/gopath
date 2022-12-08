package main

import (
	"fmt"
)

func main() {
	v := "`1234567890-=qwertyuiop[]\asdfghjkl;'zxcvbnm,./QWERTYUIOPASDFGHJKLZXCVBNM~!@#$%^&*()_+{}|:<>?"

	fmt.Println(len(v))

	fmt.Println("end-----------")
}
