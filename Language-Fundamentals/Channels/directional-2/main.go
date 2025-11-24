package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	go foo(c)

	go bar(c)
	// runtime.Gosched()
	fmt.Println("About to exit..")
}

// send
func foo(c chan<- int) {

	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)

}
