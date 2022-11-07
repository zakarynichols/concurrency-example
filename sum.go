package main

import "fmt"

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	fmt.Printf("sent %d to chan %p\n", sum, c)
	// close(c) // close channel early to force panic
}
