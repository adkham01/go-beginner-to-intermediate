package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
	wg.Wait()
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
}
