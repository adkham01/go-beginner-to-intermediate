package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("OS:", runtime.NumCPU())
	fmt.Println("OS:", runtime.GOARCH)

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("\n", counter)

}
