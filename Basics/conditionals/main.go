package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Sequence
	fmt.Println("This is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is third statement to run
	y := 5  // this is fourth statement to run
	// Conditional
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to or greater than the meaning of life")
	}
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("Greater than the meaning of life")
	}

	if x < 42 && y < 42 {
		fmt.Println("Both are less than the meaning of life")
	}
	if x > 30 || x < 42 {
		fmt.Println("x is getting close the meaning of life")
	}
	if x != 42 {
		fmt.Println("x is not meaning of life")
	}

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is lesser than or equal x which is %v\n", z, x)
	}

	switch {
	case x < 42:
		fmt.Println("x is LESS than 42")
	case x == 42:
		fmt.Println("x is EQUAL to 42")
	case x > 42:
		fmt.Println("x is GRATER than 42")
	default:
		fmt.Println("This is the default case for 42")
	}

	// no default  fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
	default:
		fmt.Println("Printed because of fallthrough: this is default case for x")
	}

	// no default  fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("Printed because of fallthrough: this is default case for x")
	}

	// concurrency
	// switch on channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int31n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Microsecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1 ", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2 ", v2)
	}

	// for loops
	for i := 5; i < 5; i++ {
		fmt.Printf("counting to five %v first for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t second for loop\n", y)
		y++
	}

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("continuing even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("----")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// for range loop\
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":    42,
		"My Penny": 32,
	}

	for k, v := range m {
		fmt.Printf("ranging over a map %v\t%v\n", k, v)
	}
}
