package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	// close(c) // close channel early to force panic
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)     // create a channel before using
	go sum(s[:len(s)/2], c) // s[:len(s)/2] === s[0:len(s)/2] === []int{7, 2, 8}
	go sum(s[len(s)/2:], c) // s[len(s)/2:] === s[len(s)/2:6] === []int{-9, 4, 0}
	x, y := <-c, <-c        // receive from c
	fmt.Println(x, y, x+y)

	const max_buffer_len = 2
	ch := make(chan int, max_buffer_len) // initialize with a max buffer length
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // force a deadlock
}
