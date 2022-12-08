package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `name`
	Age   uint64 `age`
	Child *User  `child`
}

func fct(s string, fc func(str string) string) {
	if fc != nil {
		fmt.Println(fc(s))
	}
}

func printnow(s string) string {
	return s
}

func main() {
	js := `[{"name":"zs","age":18,"child":{"name":"zs","age":18}}]`
	users := []User{}

	json.Unmarshal([]byte(js), &users)

	fmt.Println(users)

	vs, _ := json.Marshal(users)
	fmt.Println(string(vs))

	fct("skr", printnow)
}
