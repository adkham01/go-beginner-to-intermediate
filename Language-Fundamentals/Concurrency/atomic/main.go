package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}
