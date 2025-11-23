package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (m mathError) Error() string {
	return fmt.Sprintf("Math error occuerd : %v %v %v", m.lat, m.long, m.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		name := fmt.Errorf("Math errror sqrt of negative")
		return 0, mathError{"50.2289 N", "99.4556 W", name}
	}
	return 42, nil
}
