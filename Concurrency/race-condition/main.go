package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			v := counter
			//time.Sleep(time.Millisecond * 100)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}
