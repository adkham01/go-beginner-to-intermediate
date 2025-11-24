package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("About to exit...")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := range 100 {
			c <- i + 1
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
