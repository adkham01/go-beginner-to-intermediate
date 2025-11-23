package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error happened", err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error happened", err)
		// log.Println("Error happened", err)
		// log.Fatal(err)
		// panic(err)
	}
	defer f2.Close()
	log.SetOutput(f2)

	fmt.Println("About to exit check log.txt")
}
