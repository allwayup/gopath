package main

import (
	"encoding/json"
	"fmt"
)

type Huffuman struct {
	Char  int
	Count int
	Left  *Huffuman
	Right *Huffuman
}

func toHuffuman(s string) (map[int]*Huffuman, []*Huffuman) {
	r := make(map[int]*Huffuman)
	l := len(s)
	a := make([]*Huffuman, l)
	i := 0

	for _, v := range s {
		vi := int(v)
		c := r[vi]
		if c == nil {
			c = &Huffuman{
				Char:  vi,
				Count: 0,
			}
			r[vi] = c
		}
		a[i] = c
		c.Count = c.Count + 1
		i++
	}
	return r, a
}

func main() {
	s := "法外狂徒张三gank罗翔被反杀"
	r, a := toHuffuman(s)

	s1, _ := json.Marshal(r)
	s2, _ := json.Marshal(a)

	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println("end..............")
}
