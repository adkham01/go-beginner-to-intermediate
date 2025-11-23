package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("About to exit...")
}

// populate channel
func populate(c chan int) {
	for i := range 10 {
		c <- i

	}
	close(c)
}

// fanOutIn channel
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const gt = 10

	for range gt {

		wg.Go(func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
		})
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn((1000))
}
