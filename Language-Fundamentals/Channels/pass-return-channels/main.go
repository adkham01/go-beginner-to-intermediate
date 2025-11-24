package main

import "fmt"

func main() {

	for n := range puller(incrementer()) {
		fmt.Println(n)
	}

	c := make(chan int)
	go func() {
		fmt.Println("Calculating factorial...")
		c <- 5
	}()
	for f := range factorial(c) {
		fmt.Println(f)
	}
}

func incrementer() chan int {
	out := make(chan int)
	go func() {
		for i := range 10 {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		sum := 1
		start := <-c
		for i := start; i > 0; i-- {
			sum *= i
		}
		out <- sum
		close(out)
	}()
	return out
}
