package main

import (
	"fmt"
)

type HuffumanFlag struct {
	Flag int
	Prev *HuffumanFlag
	Next *HuffumanFlag
}

func Sort() {
	end := &HuffumanFlag{Flag: 1}
	start := &HuffumanFlag{Flag: 1}
	c = &HuffumanFlag{Flag: 1, Prev: start, Next: end}
	start.Next = c
	end.Prev = c

	var p *HuffumanFlag
	if c.Prev != nil && c.Flag > c.Prev.Flag {
		p = c.Prev.Prev
		if c.Next == nil {
			end = c.Prev
		} else {
			c.Next.Prev = c.Prev
		}
		c.Prev.Next = c.Next
		c.Prev = c.Prev.Prev
	} else {
		break
	}
	for {
		if c.Prev != nil && c.Flag > c.Prev.Flag {
			p = c.Prev.Prev
			c.Prev = c.Prev.Prev
		} else {
			break
		}
	}
	if p != nil {
		p.Next = c
	}
	c.Prev.Prev = c
	c.Prev.Next = c.Next
	if c.Next != nil {
		c.Next.Prev = c.Prev
	}
	c.Next = c.Prev
	c.Prev = p
	if c.Prev == nil {
		start = c
	}
}

func backup() {
	// for {
	// 	if c.Prev != nil && c.Flag > c.Prev.Flag {
	// 		p := c.Prev.Prev
	// 		if p == nil {
	// 			start = c
	// 		} else {
	// 			p.Next = c
	// 		}
	// 		c.Prev.Prev = c
	// 		c.Prev.Next = c.Next
	// 		if c.Next == nil {
	// 			end = c.Prev
	// 		} else {
	// 			c.Next.Prev = c.Prev
	// 		}
	// 		c.Next = c.Prev
	// 		c.Prev = p
	// 	} else {
	// 		break
	// 	}
	// }
}

func main() {
	Sort()
	fmt.Println("end.........................")
}
