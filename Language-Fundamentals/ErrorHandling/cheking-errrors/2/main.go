package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")

	if err != nil {
		fmt.Println("Error", err)
	}

	defer f.Close()

	r := strings.NewReader("Whats Up")
	io.Copy(f, r)
}
