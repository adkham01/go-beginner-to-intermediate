package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(bs))
}
