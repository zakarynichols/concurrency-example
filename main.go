package main

import (
	"fmt"
)

func main() {
	// goroutines
	strChan := make(chan string)
	go Say("world", strChan)
	go Say("hello", strChan)

	sayX, sayY := <-strChan, <-strChan
	fmt.Println(sayX, sayY)

	var intChan chan int // reuse chan declaration

	// channels
	s := []int{7, 2, 8, -9, 4, 0}
	intChan = make(chan int)      // create a channel before using
	go Sum(s[:len(s)/2], intChan) // s[:len(s)/2] === s[0:len(s)/2] === []int{7, 2, 8}
	go Sum(s[len(s)/2:], intChan) // s[len(s)/2:] === s[len(s)/2:6] === []int{-9, 4, 0}
	x, y := <-intChan, <-intChan  // receive from c
	fmt.Println(x, y, x+y)

	// buffered channels
	const max_buffer_len = 2
	intChan = make(chan int, max_buffer_len) // initialize with a max buffer length
	intChan <- 1
	intChan <- 2
	// fmt.Println(<-ch) // force a deadlock

	// fibonnaci
	intChan = make(chan int, 10)
	go Fibonacci(cap(intChan), intChan)
	for i := range intChan {
		fmt.Println(i)
	}
}
