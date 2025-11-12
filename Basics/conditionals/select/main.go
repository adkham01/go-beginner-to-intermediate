package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int31n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Microsecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1 ", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2 ", v2)
	}
}
