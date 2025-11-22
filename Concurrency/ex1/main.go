package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
	fmt.Println("About to exit")

}

func foo() {
	for i := range 10 {
		fmt.Println("Func foo", i)
	}
	wg.Done()
}

func bar() {
	for i := range 10 {
		fmt.Println("Func bar", i)
	}
	wg.Done()
}
