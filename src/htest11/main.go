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

func printJson(va []*HuffumanFlag) {
	for i := range va {
		v := va[i]
		fmt.Println(v)
	}
}

func one() {
	a := &HuffumanFlag{Flag: 1}
	b := &HuffumanFlag{Flag: 2}
	c := &HuffumanFlag{Flag: 3}
	d := &HuffumanFlag{Flag: 4}
	a.Next = b
	b.Prev = a
	b.Next = c
	c.Prev = b
	c.Next = d
	d.Prev = c
	printJson([]*HuffumanFlag{a, b, c, d})

	c.Prev = b.Prev
	c.Prev.Next = c
	b.Prev = c
	b.Next = c.Next
	c.Next = b
	printJson([]*HuffumanFlag{a, b, c, d})
}

func two() {
	start := &HuffumanFlag{Flag: 1}
	end := &HuffumanFlag{Flag: 2}
	start.Next = end
	end.Prev = start
	d := start
	c := end

	printJson([]*HuffumanFlag{start, end, c, d})

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

	fmt.Println()

	printJson([]*HuffumanFlag{start, end, c, d})
}

func main() {
	lock := new(sync.Mutex)
	lock.Lock()
	two()
	lock.Unlock()
}
