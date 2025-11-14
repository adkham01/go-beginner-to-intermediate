package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		model: "M5",
		make:  "BMW",
		doors: 4,
		color: "black",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		model: "Tesla Model 3",
		make:  "Tesla",
		doors: 4,
		color: "white",
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("--------------")

	fmt.Println(v1.engine.electric)
	fmt.Println(v2.engine.electric)

}
