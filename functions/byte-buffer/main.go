package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.firstName))
	return err
}

func main() {
	p := person{
		firstName: "Tom",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	fmt.Println("Writing has started")
	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String(), "is stored in the buffer")
	fmt.Println("Writing has ended")
}
