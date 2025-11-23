package main

import (
	"fmt"
	"os"
)

// Println formats using the default formats for its operand and writes
// standard output
func main() {
	_, err := os.Open("no-file-txt")
	if err != nil {
		fmt.Println("Error happened ", err)
	}
}
