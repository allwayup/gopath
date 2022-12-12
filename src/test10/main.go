package main

import (
	"fmt"
	"net/http"
)

func req01() {
	response, err := http.Get("http://baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("req01:")
	fmt.Println(response)
}

func req02() {
	request, err := http.NewRequest("GET", "http://baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := http.DefaultClient.Do(request)
	fmt.Print("req02:")
	fmt.Println(response)
}

func main() {
	req01()
	req02()
	fmt.Println("end..................")
}
