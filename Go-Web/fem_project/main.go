package main

import (
	"fmt"

	"github.com/adkham01/fem_project/internal/app"
)

func main() {
	fmt.Println("Hello World")
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("Starting Application...")

}
