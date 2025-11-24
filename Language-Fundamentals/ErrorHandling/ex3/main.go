package main

import (
	"fmt"
)

type MyError struct {
	message string
}

func (m MyError) Error() string {
	return fmt.Sprintf("There is an error occured %v", m.message)
}

func main() {
	me := MyError{
		message: "Need a break",
	}

	foo(me)
}

func foo(e error) {
	fmt.Println("Foo ran err -> ", e.(MyError).message)
}
