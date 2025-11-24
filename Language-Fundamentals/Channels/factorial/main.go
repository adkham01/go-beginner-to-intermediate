package main

import "fmt"

func main() {
	in := gen()
	f := factorial(in)
	for v := range f {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			for j := 1; j <= 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
