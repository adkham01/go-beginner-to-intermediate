package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := user{
		First: "Jane",
		Last:  "Tomas",
		Age:   35,
		Sayings: []string{
			"Drop by drop bucket gets full",
			"Persistently patiently you are bounce to succeed",
		},
	}

	u3 := user{
		First: "Oliver",
		Last:  "Girud",
		Age:   35,
		Sayings: []string{
			"Success cames by sacrifise",
			"Start from NOW!",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did somthing wrong", err)
	}
}
