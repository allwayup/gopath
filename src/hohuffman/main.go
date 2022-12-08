package main

import (
	"fmt"
	"sync"
)

type HuffumanFlag struct {
	Flag int
	Prev *HuffumanFlag
	Next *HuffumanFlag
}

func CountChar(s string) *HuffumanFlag {
	r := []rune(s)

	h := make(map[int32]*HuffumanFlag)

	end := &HuffumanFlag{Flag: 1}
	start := &HuffumanFlag{Flag: 1, Next: end}
	end.Prev = start

	// 升序排列
	f := r[0]
	o := r[1]
	if f > o {
		h[f] = start
		h[o] = end
	} else {
		h[f] = end
		h[o] = start
	}

	// 1.统计单字符出现的次数;2.并排序
	l := len(r) - 2
	for i := 2; i < l; i++ {
		c := h[r[i]]
		if c == nil {
			c = &HuffumanFlag{Flag: 1}
			c.Prev = end
			end.Next = c
			end = c
			h[r[i]] = c
		} else {
			c.Flag++
			lock := new(sync.Mutex)
			lock.Lock()
			for {
				if c.Prev != nil && c.Flag > c.Prev.Flag {
					p := c.Prev.Prev
					if p == nil {
						start = c
					} else {
						p.Next = c
					}
					c.Prev.Prev = c
					c.Prev.Next = c.Next
					if c.Next == nil {
						end = c.Prev
					} else {
						c.Next.Prev = c.Prev
					}
					c.Next = c.Prev
					c.Prev = p
				} else {
					break
				}
			}
			lock.Unlock()
		}
	}
	return start
}

func printLinkList(v *HuffumanFlag) {
	n := v
	for {
		fmt.Printf("%p: ", n)
		fmt.Println(n)
		if n.Next == nil {
			break
		}
		n = n.Next
	}
}

func main() {
	s := "-法外狂徒张三gank罗翔被反杀,-----罗翔gank张三33"
	h := CountChar(s)
	printLinkList(h)
}
