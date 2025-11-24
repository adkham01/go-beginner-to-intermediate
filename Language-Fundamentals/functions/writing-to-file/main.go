package main

import (
	"fmt"
	"log"
	"os"
)

// type person struct {
// 	firstName string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.firstName))
// 	return err
// }

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	s := []byte("Hello gopers")

	fmt.Println("Writing has started")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s ", err)
	}

	fmt.Println("Writing has finished")
}
