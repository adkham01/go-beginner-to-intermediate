package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for range gs {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println(counter)
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("\n", counter)

}
