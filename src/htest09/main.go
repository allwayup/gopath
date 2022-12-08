package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println(reflect.TypeOf(inputReader))
}
