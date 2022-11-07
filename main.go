package main

import (
	"fmt"
)

func main() {
	var c chan int // reuse chan declaration

	// goroutines
	s := []int{7, 2, 8, -9, 4, 0}
	c = make(chan int)      // create a channel before using
	go Sum(s[:len(s)/2], c) // s[:len(s)/2] === s[0:len(s)/2] === []int{7, 2, 8}
	go Sum(s[len(s)/2:], c) // s[len(s)/2:] === s[len(s)/2:6] === []int{-9, 4, 0}
	x, y := <-c, <-c        // receive from c
	fmt.Println(x, y, x+y)

	// buffered channels
	const max_buffer_len = 2
	c = make(chan int, max_buffer_len) // initialize with a max buffer length
	c <- 1
	c <- 2
	// fmt.Println(<-ch) // force a deadlock

	// fibonnaci
	c = make(chan int, 10)
	go Fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
