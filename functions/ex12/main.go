package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 5_000; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
