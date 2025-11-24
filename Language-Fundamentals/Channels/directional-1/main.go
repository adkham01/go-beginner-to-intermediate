package main

import "fmt"

func main() {
	c := make(chan int, 2)
	cr := make(<-chan int, 2) // receive only channel
	cs := make(chan<- int, 2) // send only channel

	// go func() {
	c <- 42
	c <- 43
	// }()

	fmt.Printf("%v\t%T\n", <-c, c)
	fmt.Printf("%v\t%T\n", cr, cr)
	fmt.Printf("%v\t%T\n", cs, cs)
}
