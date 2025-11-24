package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("About to exit...")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := range 100 {
			c <- i + 1
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return

		}
	}
}
