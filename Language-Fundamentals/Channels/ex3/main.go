package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(10)
	put(c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	pull(c)

	fmt.Println("About to exit...")
}

func put(c chan int, wg *sync.WaitGroup) {
	for range 10 {
		go func() {
			defer wg.Done()
			for i := range 10 {
				c <- i
			}
		}()
	}

}

func pull(c chan int) {
	for i := range 100 {
		fmt.Println(i+1, <-c)
	}
}
