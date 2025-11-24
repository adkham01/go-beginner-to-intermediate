package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) {
	log.Println("Log from Wrapper", s.String())
}

func main() {
	b := book{
		title: "Harry Potter",
	}
	var c count = 42

	log.Println(b)
	log.Println(c)
	println("--------------------------")
	logInfo(b)
	logInfo(c)
}

// func (r receiver) identifier(p params) (r returns){...}
