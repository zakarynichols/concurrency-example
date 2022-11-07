package main

import (
	"time"
)

func Say(s string, c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	c <- s
}
