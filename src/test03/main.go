package main

import (
	"fmt"
)

type People interface {
	setName(n string)
}

type User struct {
	name string
	age  uint64
}

func (u *User) setName(n string) {
	u.name = n
}

func main() {
	fmt.Println("start....")
	user := User{
		name: "zs",
		age:  18,
	}
	user.setName("-----")
	fmt.Println(user)
}
