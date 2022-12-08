package main

import (
	"bytes"
	"container/heap"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	str := `
	{
	   "nick_name": "Lucy",
	   "user_age": 28
	}`

	var buf bytes.Buffer
	_ = json.Compact(&buf, []byte(str))
	buf.WriteTo(os.Stdout)
	fmt.Println()

}
