package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("About to exit..")
}

// send
func send(e, o chan<- int, q chan<- bool) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- true
	close(q)
}

// receive
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even chanel:\t", v)
		case v := <-o:
			fmt.Println("From the odd chanel:\t", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from commaok", i, ok)
			} else {
				fmt.Println("from commaok", i, ok)

			}
			return
		}

	}
}
