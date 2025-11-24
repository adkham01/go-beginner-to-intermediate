package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Harry Potter",
	}
	var c count = 42

	fmt.Println(b)
	fmt.Println(c)
}

// func (r receiver) identifier(p params) (r returns){...}
