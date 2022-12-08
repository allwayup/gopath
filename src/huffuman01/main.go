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
	if len(s) <= 0 {
		return nil
	}
	r := []rune(s)

	h := make(map[int32]*HuffumanFlag)

	// 升序排列
	var start, end *HuffumanFlag
	l := len(r)
	i := 0
	for ; i < l; i++ {
		c := h[r[i]]
		if c == nil {
			c = &HuffumanFlag{Flag: 1}
			h[r[i]] = c
			if start == nil {
				start = c
			} else {
				end = c
				i++
				break
			}
		} else {
			c.Flag++
		}
	}
	if end == nil {
		return start
	}
	if start.Flag > end.Flag {
		start.Next = end
		end.Prev = start
	} else {
		start.Prev = end
		end.Next = start
		start = end
		end = end.Next
	}
	if i >= l {
		return start
	}
	// 1.统计单字符出现的次数;2.并排序
	for ; i < l; i++ {
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
			if c.Prev != nil && c.Flag > c.Prev.Flag {
				c.Prev.Next = c.Next
				if c.Next == nil {
					end = c.Prev
				} else {
					c.Next.Prev = c.Prev
				}
				if c.Prev.Prev == nil {
					start = c
					c.Next = c.Prev
					c.Prev = nil
					c.Next.Prev = c
					continue
				} else {
					c.Prev = c.Prev.Prev
				}
			} else {
				continue
			}
			for {
				if c.Flag > c.Prev.Flag {
					if c.Prev.Prev == nil {
						start = c
						c.Next = c.Prev
						c.Prev = nil
						c.Next.Prev = c
						break
					} else {
						c.Prev = c.Prev.Prev
					}
				} else {
					c.Next = c.Prev.Next
					c.Prev.Next = c
					c.Next.Prev = c
					break
				}
			}
			lock.Unlock()
		}
	}
	return start
}

func printLinkList(v *HuffumanFlag) {
	if v == nil {
		return
	}
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
	s := "-法外狂徒张三gank罗翔被反杀,-----罗翔gank张三333"
	// s := "-"
	// s := "-----"
	// s := "-杀杀杀杀杀杀杀"
	// s := "杀杀杀杀杀杀杀-"
	// s := ""
	// var s string
	h := CountChar(s)
	printLinkList(h)
}
