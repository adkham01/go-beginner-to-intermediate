package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("About to exit..")
}

// send
func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}

// receive
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even chanel: ", v)
		case v := <-o:
			fmt.Println("From the odd chanel: ", v)
		case v := <-q:
			fmt.Println("From the quit chanel: ", v)
			return
		}

	}
}
