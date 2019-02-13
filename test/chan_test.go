package main

import (
	"fmt"
	"os"
	"testing"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestChan(t *testing.T) {
	c := make(chan int, 100)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(os.Getenv("ZOOKEEPER_SERVERS"))
}
